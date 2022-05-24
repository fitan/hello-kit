package app

import (
	"bytes"
	"context"
	"entgo.io/ent/dialect/sql"
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
	"github.com/hashicorp/consul/api"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "hello/docs"
	"hello/pkg/debug"
	"hello/pkg/ent"
	"hello/pkg/middleware"
	"hello/pkg/repository"
	"hello/pkg/services"
	"hello/utils/conf"
	"hello/utils/log"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var logger *zap.SugaredLogger

type App struct {
	r          *gin.Engine
	repository *repository.Repository
	InitAuditMid
	//services   *services.Services
	httpHandler *services.HttpHandler
	g           *run.Group
	conf        *conf.MyConf
	log         *zap.SugaredLogger
	tp          *sdktrace.TracerProvider
	db          *gorm.DB
	ent         *ent.Client
	pyroscope   *profiler.Profiler
	casbin      *casbin.Enforcer
	InitCancelInterrupt
	InitMetricsEndpoint
	//InitHttpHandler
	InitMicro
}

func (a *App) Run() error {
	return a.g.Run()
}

func RunApp(confName string) {
	//fs.Parse(os.Args[1:])

	r := gin.Default()
	opts := ginprom.NewDefaultOpts()
	opts.EndpointLabelMappingFn = func(c *gin.Context) string {
		return c.FullPath()
	}
	r.Use(ginprom.PromMiddleware(opts))
	r.Use(cors.New(cors.Config{
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "X-Total-Count"},
		ExposeHeaders: []string{"X-Total-Count"},
		//AllowOrigins: []string{"http://localhost:8080","http://localhost:3000"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	g := &run.Group{}
	app, err := InitApp(r, g, ConfName(confName))
	if err != nil {
		fmt.Printf("initapp error: %s", err.Error())
		os.Exit(1)

	}
	initAuditMid(r, app.repository)
	logger.Errorw("exit", app.Run().Error())
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

func initAuditMid(r *gin.Engine, repository *repository.Repository) InitAuditMid {
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
		if err = client.Schema.Create(context.Background()); err != nil {
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

func initMetricsEndpoint(g *run.Group, conf *conf.MyConf, debugSwitch *debug.DebugSwitch) InitMetricsEndpoint {
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

	//http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", conf.Debug.Addr)
	if err != nil {
		logger.Errorw("net.Listen", "transport", "debug/HTTP", "during", "Listen", "err", err.Error())
	}
	g.Add(func() error {
		logger.Infow("init metrics", "transport", "debug/HTTP", "addr", conf.Debug.Addr)
		return r.RunListener(debugListener)
	}, func(error) {
		debugListener.Close()
	})
	return InitMetricsEndpoint{}
}

func initDebugSwitch() *debug.DebugSwitch {
	return debug.NewDebugSwitch()
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

func initHttpServerOption(debugSwitch *debug.DebugSwitch) []ginkHttp.ServerOption {
	so := make([]ginkHttp.ServerOption, 0)
	so = append(so, ginkHttp.ServerBefore(ginkHttp.PopulateRequestContext, debug.DebugSwitchRequestContext(debugSwitch)))
	return so
}

type InitMicro struct {
}

func initMicro(g *run.Group, r *gin.Engine, conf *conf.MyConf) (InitMicro, error) {
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
			logger.Errorw("iniMicro.service.Run", "err", err.Error())
			server.Stop()
		})
	return InitMicro{}, nil
}

func initCasbin(conf *conf.MyConf) (*casbin.Enforcer, error) {
	a, err := entadapter.NewAdapter("mysql", conf.Mysql.Url)
	if err != nil {
		return nil, err
	}
	return casbin.NewEnforcer("./conf/rbac_model.conf", a)
}
