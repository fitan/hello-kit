package app

import (
	"bytes"
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	microConsul "github.com/asim/go-micro/plugins/registry/consul/v4"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/casbin/casbin/v2"
	entadapter "github.com/casbin/ent-adapter"
	"github.com/chenjiandongx/ginprom"
	ginkHttp "github.com/fitan/gink/transport/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	log2 "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/consul/api"
	natsServer "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/util/addr"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "hello/docs"
	"hello/pkg/ent"
	resource2 "hello/pkg/ent/resource"
	"hello/pkg/middleware"
	"hello/pkg/repository"
	"hello/pkg/services"
	casbin2 "hello/pkg/services/casbin"
	"hello/utils/conf"
	debug2 "hello/utils/debug"
	"hello/utils/httpclient"
	"hello/utils/log"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var logger *zap.SugaredLogger

type App struct {
	r          *gin.RouterGroup
	repository *repository.Repository
	InitAuditMid
	nats        *nats.Conn
	services    *services.Services
	debug       *debug2.DebugSwitch
	httpHandler *services.HttpHandler
	g           *run.Group
	conf        *conf.MyConf
	log         *zap.SugaredLogger
	tp          *sdktrace.TracerProvider
	db          *gorm.DB
	ent         *ent.Client
	pyroscope   *profiler.Profiler
	casbin      *casbin.SyncedEnforcer
	httpNewFunc httpclient.NewFunc
	InitCancelInterrupt
	InitMetricsEndpoint
	//InitHttpHandler
	InitMicro
}

func (a *App) Run() error {
	return a.g.Run()
}

func RunApp(confName string) {
	r := gin.Default()
	r.Use(otelgin.Middleware("hello-kit"))

	group := r.Group("/kit")
	opts := ginprom.NewDefaultOpts()
	opts.EndpointLabelMappingFn = func(c *gin.Context) string {
		return c.FullPath()
	}
	group.Use(ginprom.PromMiddleware(opts))
	group.Use(cors.New(cors.Config{
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "X-Total-Count"},
		ExposeHeaders: []string{"X-Total-Count"},
		//AllowOrigins: []string{"http://localhost:8080","http://localhost:3000"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	g := &run.Group{}
	app, err := InitApp(r, group, g, ConfName(confName))

	if err != nil {
		fmt.Printf("initapp error: %s", err.Error())
		os.Exit(1)

	}
	initAuditMid(group, app.repository)
	logger.Errorw("exit", app.Run().Error())
}

func CasbinInput(confName string) error {
	g := &run.Group{}
	r := gin.Default()
	group := r.Group("/")
	app, err := InitApp(r, group, g, ConfName(confName))
	if err != nil {
		return err
	}
	ctx := context.Background()
	app.services.Casbin.BindPermission(ctx, casbin2.Permission{
		UserId: 1,
		RoleId: 1,
		Domain: "/1/1",
	})
	app.services.Casbin.AddRoleAuthorization(ctx, casbin2.Role{
		RoleId:     1,
		ResourceId: 100,
	})
	return nil
}

func PathPutInStorage(confName string) error {
	g := &run.Group{}
	r := gin.Default()
	group := r.Group("/")
	app, err := InitApp(r, group, g, ConfName(confName))
	if err != nil {
		return err
	}
	list := app.debug.List()
	req := make([]ent.ResourceBaseCreateReq, 0)
	t := resource2.TypeAPI
	for _, v := range list {
		_, err := app.repository.Resource.ByQueriesOne(
			context.Background(), struct {
				ent.ResourceTableActionEQForm
				ent.ResourceTablePathEQForm
				ent.ResourceTableTypeEQForm
			}{
				ResourceTableActionEQForm: ent.ResourceTableActionEQForm{ActionEQ: &v.Method},
				ResourceTablePathEQForm:   ent.ResourceTablePathEQForm{PathEQ: &v.Path},
				ResourceTableTypeEQForm:   ent.ResourceTableTypeEQForm{TypeEQ: &t},
			},
		)
		if err == nil {
			continue
		}

		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}

		req = append(req, ent.ResourceBaseCreateReq{
			Name:     v.Annotation,
			Key:      path.Join(v.Method, v.Path),
			Path:     v.Path,
			Action:   v.Method,
			Comments: v.Annotation,
			Type:     resource2.TypeAPI,
		})
	}
	_, err = app.repository.Resource.CreateMany(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

type InitAuditMid struct {
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func initAuditMid(r *gin.RouterGroup, repository *repository.Repository) InitAuditMid {
	r.Use(
		func(c *gin.Context) {

			startTime := time.Now()
			blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			c.Writer = blw

			bodyB, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyB))

			c.Next()

			method := c.Request.Method
			url := c.FullPath()
			query := c.Request.URL.Query().Encode()
			remoteIP := c.Request.RemoteAddr
			response := blw.body.String()
			statusCode := c.Writer.Status()

			_, _ = repository.Audit.Create(c.Request.Context(), ent.AuditBaseCreateReq{
				URL:        url,
				Query:      query,
				Method:     method,
				Request:    string(bodyB),
				Response:   string(response),
				Header:     "",
				StatusCode: statusCode,
				RemoteIP:   remoteIP,
				ClientIP:   c.ClientIP(),
				CostTime:   fmt.Sprintf("%v", time.Now().Sub(startTime)),
			})

		})
	return InitAuditMid{}
}

func initEnt(conf *conf.MyConf) (*ent.Client, error) {
	drv, err := sql.Open("mysql", conf.Mysql.Url)
	if err != nil {
		return nil, err
	}
	client := ent.NewClient(ent.Driver(drv)).Debug()
	if conf.Mysql.AutoCreate {
		if err = client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
			return nil, err
		}
	}

	return ent.NewClient(ent.Driver(drv)), nil
}

func initDb(conf *conf.MyConf) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(conf.Mysql.Url))
}

type ConfName string

func initLog(conf *conf.MyConf) *zap.SugaredLogger {
	core := log.DefaultZapCore(conf.Log.FileName, conf.Log.Dir, zapcore.Level(conf.Log.Lervel))
	logger = zap.New(core, zap.AddCaller()).Sugar()
	return logger
}

func initConf(confName ConfName) (*conf.MyConf, error) {
	myConf := conf.MyConf{}
	_, err := conf.WatchFile(string(confName)+".conf", []string{"./conf"}, &myConf, 5*time.Second)
	if err != nil {
		return nil, err
	}
	return &myConf, nil
}

func initPyroscope(conf *conf.MyConf) (*profiler.Profiler, error) {
	if conf.Pyroscope.Open {
		return profiler.Start(
			profiler.Config{
				ApplicationName: conf.App.Name,

				// replace this with the address of pyroscope server
				ServerAddress: conf.Pyroscope.Url,

				// by default all profilers are enabled,
				// but you can select the ones you want to use:
				ProfileTypes: []profiler.ProfileType{
					profiler.ProfileCPU,
					profiler.ProfileAllocObjects,
					profiler.ProfileAllocSpace,
					profiler.ProfileInuseObjects,
					profiler.ProfileInuseSpace,
				},
			},
		)
	}
	return nil, nil
}

type InitHttpHandler struct {
}

func initHttpHandler(r *gin.Engine, g *run.Group, log *zap.SugaredLogger, conf *conf.MyConf) InitHttpHandler {

	httpListener, err := net.Listen("tcp", conf.App.Addr)
	if err != nil {
		log.Errorw(err.Error(), "transport", "HTTP", "during", "Listen")
	}
	g.Add(func() error {
		log.Infow("http listener", "transport", "HTTP", "addr", conf.App.Name)
		return http.Serve(httpListener, r)
	}, func(error) {
		httpListener.Close()
	})

	return InitHttpHandler{}

}

func initTracer(conf *conf.MyConf) *sdktrace.TracerProvider {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(conf.App.Name),
		),
	)
	if err != nil {
		logger.Fatalw("resource.New", "err", err.Error())
	}
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Trace.TracerProviderAddr)))
	if err != nil {
		logger.Fatalw("jaeger.New", "err", err.Error())
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tp)
	return tp
}

type InitMetricsEndpoint struct {
}

func initMetricsEndpoint(g *run.Group, conf *conf.MyConf, debugSwitch *debug2.DebugSwitch) InitMetricsEndpoint {
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/debug", func(c *gin.Context) {
		c.JSON(http.StatusOK, debugSwitch.List())
	})
	type SetDebug struct {
		Path   string `json:"path"`
		Method string `json:"method"`
		Enable bool   `json:"enable"`
	}
	r.POST("/debug", func(c *gin.Context) {
		var s SetDebug
		err := c.Bind(&s)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, debugSwitch.SetDebug(s.Path, s.Method, s.Enable))
		return
	})

	r.POST("/debug/rest", func(c *gin.Context) {
		debugSwitch.ResetDebug()
		c.JSON(http.StatusOK, "")
		return
	})

	debugListener, err := net.Listen("tcp", conf.Debug.Addr)
	if err != nil {
		logger.Errorw("net.Listen", "transport", "debug/HTTP", "during", "Listen", "err", err.Error())
	}
	g.Add(func() error {
		logger.Infow("init metrics", "transport", "debug/HTTP", "addr", conf.Debug.Addr)
		return r.RunListener(debugListener)
	}, func(error) {
		err = debugListener.Close()
		if err != nil {
			logger.Errorw("debug http: ", "err", err.Error())
		}
	})
	return InitMetricsEndpoint{}
}

func initDebugSwitch() *debug2.DebugSwitch {
	return debug2.NewDebugSwitch()
}

type InitCancelInterrupt struct {
}

func initCancelInterrupt(g *run.Group) InitCancelInterrupt {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
	return InitCancelInterrupt{}
}

func initEndpointMiddleware(services *services.Services, repository repository.Repository) []endpoint.Middleware {
	mw := make([]endpoint.Middleware, 0)
	mw = append(mw, middleware.PermissionMiddleware(services, repository))
	return mw
}

type ErrorHandler struct {
}

func (e ErrorHandler) Handle(ctx context.Context, err error) {
	logger.Errorw("error handler", "traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String(), "err", err.Error())
}

func initHttpServerOption(debugSwitch *debug2.DebugSwitch) []ginkHttp.ServerOption {
	errorEncoder := func(ctx context.Context, err error, gCtx *gin.Context) {
		gCtx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"err":     err.Error(),
			"traceId": trace.SpanFromContext(ctx).SpanContext().TraceID().String(),
		})
	}

	so := make([]ginkHttp.ServerOption, 0)
	so = append(so,
		ginkHttp.ServerErrorHandler(ErrorHandler{}),
		ginkHttp.ServerErrorEncoder(errorEncoder),
		ginkHttp.ServerBefore(ginkHttp.PopulateRequestContext, debug2.DebugSwitchRequestContext(debugSwitch), middleware.UrlRequestContext()))
	return so
}

type InitMicro struct {
}

func initMicro(g *run.Group, r *gin.Engine, conf *conf.MyConf) (res InitMicro, err error) {

	consulConf := api.DefaultConfig()
	consulConf.Address = conf.Consul.Addr
	registry := microConsul.NewRegistry(microConsul.Config(consulConf))

	srv := httpServer.NewServer(
		server.Name(conf.App.Name),
		server.Address(conf.App.Addr),
	)

	gin.SetMode(gin.ReleaseMode)

	hd := srv.NewHandler(r)
	if err := srv.Handle(hd); err != nil {
		return InitMicro{}, err
	}
	serivce := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry),
		micro.Flags(&cli.StringFlag{
			Name:  "conf",
			Value: "dev",
		}),
	)
	serivce.Init()

	g.Add(
		func() error {
			logger.Infow("micro run...", "transport", "HTTP", "addr", conf.App.Addr)
			return serivce.Run()
		}, func(err error) {
			err = server.Stop()
			if err != nil {
				logger.Errorw("iniMicro.service.Run", "err", err.Error())
			}
		})

	initSdConsul()
	return InitMicro{}, nil
}

func initCasbin(conf *conf.MyConf) (*casbin.SyncedEnforcer, error) {
	a, err := entadapter.NewAdapter("mysql", conf.Mysql.Url)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewSyncedEnforcer("./conf/rbac_model.conf", a)
	if err != nil {
		return nil, err
	}
	e.StartAutoLoadPolicy(10 * time.Second)
	return e, err
	//return casbin.NewEnforcer("./conf/rbac_model.conf", a)
}

func initDefaultHttpClient() httpclient.NewFunc {
	return func(fs ...httpclient.Option) *resty.Client {
		fs = append(fs,
			httpclient.WithDebugMid(logger, 10240),
			httpclient.WithTimeOut(30*time.Second),
		)
		return httpclient.New(fs...)
	}
}

func initNats() (res *nats.Conn, err error) {
	return nats.Connect("nats://localhost:4222", nats.Name("my-nats"))
}

//func initCache(conf *conf.MyConf) *marshaler.Marshaler {
//	promMetrics := metrics.NewPrometheus("hello")
//	redisStore := store.NewRedis(redis.NewClient(&redis.Options{Addr: conf.Redis.Addr}),nil)
//	cacheManager := cache.NewMetric(promMetrics, cache.New(redisStore))
//	return marshaler.New(cacheManager)
//}

// ???????????????consul
func initSdConsul() (res sd.Registrar, err error) {
	consulConf := api.DefaultConfig()
	consulConf.Address = "localhost:8500"
	consulClient, err := api.NewClient(consulConf)
	if err != nil {
		err = errors.Wrap(err, "consul.NewClient")
		return
	}
	ccc := consulsd.NewClient(consulClient)

	extract, err := addr.Extract("")
	if err != nil {
		return nil, err
	}

	reg := consulsd.NewRegistrar(ccc, &api.AgentServiceRegistration{
		ID:      "localhost",
		Name:    "consulSD",
		Tags:    nil,
		Port:    8080,
		Address: extract,
		Check: &api.AgentServiceCheck{
			TCP:                            extract + ":8080",
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "12s",
		},
	}, log2.NewNopLogger())
	reg.Register()

	i := consulsd.NewInstancer(ccc, log2.NewNopLogger(), "hello-kit", []string{}, true)
	c := make(chan sd.Event, 0)
	go func() {
		i.Register(c)
		for v := range c {
			logger.Infow("consulSD", "event", v.Instances)
		}
	}()
	return
}

func initNatsServer() (res *natsServer.Server, err error) {
	res, err = natsServer.NewServer(&natsServer.Options{
		Host: "localhost",
		Port: 0,
		//JetStream: true,
		//StoreDir: "./data",
		Debug: true,
	})

	if err != nil {
		panic(err)
	}

	go res.Start()
	return
}
