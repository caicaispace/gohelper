package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Grpc struct {
	GrpcServer *grpc.Server
	listener   net.Listener
}

func New(addr string) *Grpc {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	// 实例化grpc服务端
	s := grpc.NewServer()
	return &Grpc{
		GrpcServer: s,
		listener:   lis,
	}
}

func (g *Grpc) Start() {
	// 往grpc服务端注册反射服务
	reflection.Register(g.GrpcServer)
	// 启动grpc服务
	if err := g.GrpcServer.Serve(g.listener); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
