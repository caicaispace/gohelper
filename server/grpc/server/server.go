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
	addr       string
}

func NewServer(addr string) *Grpc {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	// 实例化grpc服务端
	s := grpc.NewServer()
	return &Grpc{
		GrpcServer: s,
		listener:   lis,
		addr:       addr,
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

// type Grpc struct {
// 	GrpcServer *grpc.Server
// 	listener   net.Listener
// 	addr       string
// }

// func NewServer() *Grpc {
// 	return &Grpc{}
// }

// func (g *Grpc) SetAddr(addr string) *Grpc {
// 	g.addr = addr
// 	return g
// }

// func (g *Grpc) Start() *Grpc {
// 	lis, err := net.Listen("tcp", g.addr)
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to listen: %v", err))
// 	}
// 	// 实例化grpc服务端
// 	s := grpc.NewServer()
// 	g.GrpcServer = s
// 	g.listener = lis
// 	// 往grpc服务端注册反射服务
// 	reflection.Register(g.GrpcServer)
// 	// 启动grpc服务
// 	if err := g.GrpcServer.Serve(g.listener); err != nil {
// 		panic(fmt.Sprintf("failed to serve: %v", err))
// 	}
// 	return g
// }
