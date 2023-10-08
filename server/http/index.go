package http

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/caicaispace/gohelper/logx"
	"github.com/caicaispace/gohelper/print"
	"github.com/caicaispace/gohelper/setting"
	"github.com/gin-gonic/gin"
)

var (
	// server
	serverMode = flag.String("mode", gin.DebugMode, "Server: run mode")
	count      int64
)

type IRouter interface {
	Router(server *Server)
}

type IMiddleware interface {
	Use(r *gin.Engine)
}

func init() {
	setting.Server.RunMode = *serverMode
}

type Server struct {
	*gin.Engine
	group       *gin.RouterGroup
	serverAddr  string
	serverHost  string
	serverPort  string
	openProfile bool
	beforeFunc  func(env string)
}

func NewServer() *Server {
	gin.DefaultWriter = io.Discard // disable router map log
	return &Server{
		Engine: gin.New(),
		// Engine: gin.Default(),
	}
}

func (s *Server) AddMiddlewares(middlewares ...IMiddleware) *Server {
	for _, m := range middlewares {
		m.Use(s.Engine)
	}
	return s
}

func (s *Server) AddRoutes(routers ...IRouter) *Server {
	for _, c := range routers {
		c.Router(s)
	}
	return s
}

func (s *Server) AddGroupRouters(group string, routers ...IRouter) *Server {
	s.group = s.Group(group)
	for _, c := range routers {
		c.Router(s)
	}
	return s
}

func (s *Server) AddBeforeFunc(fn func(env string)) *Server {
	s.beforeFunc = fn
	return s
}

func (s *Server) SetBeforeFunc(fn func(env string)) *Server {
	s.AddBeforeFunc(fn)
	return s
}

func (s *Server) AddServerAddr(addr string) *Server {
	addrArr := strings.Split(addr, ":")
	s.serverHost = addrArr[0]
	s.serverPort = addrArr[1]
	s.serverAddr = addr
	return s
}

func (s *Server) SetServerAddr(addr string) *Server {
	s.AddServerAddr(addr)
	return s
}

func (s *Server) SetOpenProfile(status bool) *Server {
	s.openProfile = status
	return s
}

// var f *os.File

func (s *Server) Start() {
	// //CPU 性能分析
	// runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	// runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	// runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
	// f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	l.Error(err)
	// 	return
	// }
	// pprof.StartCPUProfile(f)
	s.Engine.Use(gin.Logger())
	s.Engine.Use(gin.Recovery())
	if s.beforeFunc != nil {
		s.beforeFunc(setting.Server.Env)
	}
	if s.serverAddr != "" {
		setting.Server.Addr = s.serverAddr
	}
	s.registerDefaultRouter()
	print.CommandPrint(print.CommandSetPrintData("http", setting.Server.Addr, setting.Server.RunMode))
	maxHeaderBytes := 1 << 20
	gin.SetMode(setting.Server.RunMode)
	httpServer := &http.Server{
		Addr:           setting.Server.Addr,
		Handler:        s.Engine,
		ReadTimeout:    time.Duration(setting.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	go httpServer.ListenAndServe()
	listenSignal(context.Background(), httpServer)
}

func (s *Server) registerDefaultRouter() {
	s.Engine.GET("/check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok "+fmt.Sprint(count)+" remote:"+c.Request.RemoteAddr+" "+c.Request.URL.String())
		count++
	})
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-sc
	logx.Infof("exit: signal=<%d>.", sig)
	switch sig {
	case syscall.SIGTERM:
		logx.Infof("exit: bye :).")
		os.Exit(0)
	default:
		logx.Infof("exit: bye :(.")
		// // CPU 性能分析
		// f.Close()
		// pprof.StopCPUProfile()
		os.Exit(1)
	}
}
