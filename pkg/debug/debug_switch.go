package debug

import (
	"context"
	"fmt"
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	"sync"
)

func DebugSwitchRequestContext(ds *DebugSwitch) func(ctx context.Context, r *gin.Context) context.Context {
	return func(ctx context.Context, r *gin.Context) context.Context {
		has, _ := ds.Debug(r.FullPath(), r.Request.Method)
		return context.WithValue(ctx, http.ContextKeyRequestDebug, has)
	}
}

type contextKey int

const (
	ContextKeyDebugEnable contextKey = iota
)

type DebugSwitch struct {
	l    []msg
	m    map[string]bool
	lock sync.RWMutex
}

func NewDebugSwitch() *DebugSwitch {
	return &DebugSwitch{
		l: make([]msg, 0),
		m: make(map[string]bool, 0),
	}
}

type msg struct {
	Annotation string
	Path       string
	Method     string
	Enable     bool
}

func (d *DebugSwitch) Register(annotation, path, method string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.m[path+method] = false
	d.l = append(d.l, msg{
		Annotation: annotation,
		Path:       path,
		Method:     method,
	})
}

func (d *DebugSwitch) SetDebug(path, method string, b bool) (err error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.m[path+method]
	if ok {
		d.m[path+method] = b
		return nil
	}

	return fmt.Errorf("not locate path and method")
}

func (d *DebugSwitch) ResetDebug() {
	d.lock.Lock()
	defer d.lock.Unlock()
	for index, _ := range d.m {
		d.m[index] = false
	}
	return
}

func (d *DebugSwitch) Debug(path, method string) (bool, error) {
	has, ok := d.m[path+method]
	if ok {
		return has, nil
	}
	return false, fmt.Errorf("not locate path and method")

}

func (d *DebugSwitch) List() []msg {
	for i, _ := range d.l {
		d.l[i].Enable = d.m[d.l[i].Path+d.l[i].Method]
	}
	return d.l
}
