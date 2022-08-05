package _example

import (
	httpServer "github.com/caicaispace/gohelper/server/http"
	"github.com/caicaispace/gohelper/server/http/_example/middleware"
	"github.com/gin-gonic/gin"
)

const serverAddr = "127.0.0.1:9601"

func NewServer() {
	s := httpServer.NewServer()
	s.AddServerAddr(serverAddr)
	s.AddMiddlewares(
		// middleware.NewTrace(),
		middleware.NewGrafana(),
	)
	s.AddRoutes(
		NewTestHandle(),
	)
	// s.AddGroupRouters("/v1/api",
	// 	NewTestHandle(),
	// )
	s.Start()
}

type TestHandle struct {
	serviceName     string
	routerGroupPath string
}

func NewTestHandle() *TestHandle {
	return &TestHandle{
		serviceName:     "test",
		routerGroupPath: "/v1/api",
	}
}

func (th *TestHandle) Router(server *httpServer.Server) {
	router := server.Engine.Group(th.routerGroupPath)
	{
		router.GET("/test", th.Test())
		router.GET("/test_pager", th.TestPager())
	}
	// server.Handle(http.MethodGet, "/v1/api/test", th.Test())
	// server.Handle(http.MethodGet, "/v1/api/test_pager", th.TestPager())
}

func (th *TestHandle) Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := httpServer.Context{C: c}
		ctx.Success(nil, nil)
	}
}

func (th *TestHandle) TestPager() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := httpServer.Context{C: c}
		pager := ctx.GetPager()
		pager.SetTotal(100)
		ctx.Success(gin.H{
			"page":  pager.GetPage(),
			"limit": pager.GetLimit(),
			"total": pager.GetTotal(),
		}, nil)
	}
}
