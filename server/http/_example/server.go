package _example

import (
	"github.com/caicaispace/gohelper/server/http"
	httpServer "github.com/caicaispace/gohelper/server/http"
	"github.com/gin-gonic/gin"
)

const serverAddr = "127.0.0.1:9601"

func NewServer() {
	s := httpServer.NewServer()
	s.SetServerAddr(serverAddr)
	apiV1 := s.Engine.Group("/v1/api")
	{
		apiV1.GET("/test", Test)
		apiV1.GET("/test_pager", TestPager)
	}
	s.Start()
}

func Test(c *gin.Context) {
	ctx := http.Context{C: c}
	ctx.Success(nil, nil)
}

func TestPager(c *gin.Context) {
	ctx := http.Context{C: c}
	pager := ctx.GetPager()
	pager.SetTotal(100)
	ctx.Success(gin.H{
		"page":  pager.GetPage(),
		"limit": pager.GetLimit(),
		"total": pager.GetTotal(),
	}, nil)
}
