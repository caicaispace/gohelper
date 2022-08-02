package server

import (
	"context"
	"fmt"
	"time"

	"github.com/caicaispace/gohelper/cluster/etcd"
	"github.com/caicaispace/gohelper/server/grpc/_example"
	"github.com/caicaispace/gohelper/server/grpc/server"
	clientV3 "go.etcd.io/etcd/client/v3"
)

const serverAddr = "127.0.0.1:9601"
const serviceName = "hello"

func NewServer() {
	register, _ := etcd.NewRegister(&etcd.NodeInfo{
		Addr:     serverAddr,
		Name:     serviceName,
		UniqueId: fmt.Sprintf("discovery/%s/instance_id/%s", serviceName, "888"),
	}, clientV3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		DialTimeout:          2 * time.Second,
		DialKeepAliveTime:    time.Second,
		DialKeepAliveTimeout: time.Second,
	})
	go register.Run()
	s := server.NewServer().SetAddr(serverAddr).Start()
	_example.RegisterHelloWorldServer(s.GrpcServer, &HelloWorldService{})
}

type HelloWorldService struct{}

func (h *HelloWorldService) SayHello(ctx context.Context, in *_example.SayHelloReq) (*_example.SayHelloRsp, error) {
	return &_example.SayHelloRsp{Message: "Hello " + in.Name}, nil
}
