package _example

import (
	"context"
	"fmt"
	"time"

	"github.com/caicaispace/gohelper/cluster/etcd"
	"github.com/caicaispace/gohelper/server/grpc/_example/hello"
	"github.com/caicaispace/gohelper/server/grpc/client"
	clientV3 "go.etcd.io/etcd/client/v3"
)

const timeFormat = "15:04:05"
const serviceName = "hello"

func NewClient() {
	manager := etcd.NewNodeManager()

	dis, _ := etcd.NewDiscovery(&etcd.NodeInfo{}, clientV3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		DialTimeout:          2 * time.Second,
		DialKeepAliveTime:    time.Second,
		DialKeepAliveTimeout: time.Second,
	}, manager)

	dis.Pull()
	go dis.Watch()

	c := client.GetInstance()
	for _, v := range manager.All() {
		for nodeId, node := range v {
			c.AddConn(node.Name, node.Addr, nodeId)
		}
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		fmt.Println("")
		fmt.Println("---------------------------------------------")
		fmt.Println("")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		c2 := hello.NewHelloClient(client.GetInstance().Pick(serviceName))
		resp, err := c2.SayHello(ctx, &hello.SayHelloReq{
			Name: "caicaispace",
		})
		if err != nil {
			fmt.Printf("warning ⛔ %s X %s\n", time.Now().Format(timeFormat), err.Error())
		} else {
			fmt.Println("")
			fmt.Printf("I'm client 👉 %s => %s\n\n", time.Now().Format(timeFormat), resp.Message)
		}
		cancel()
	}
}
