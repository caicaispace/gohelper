package _example_test

import (
	"testing"
	"time"

	"github.com/caicaispace/gohelper/server/grpc/_example/client"
	"github.com/caicaispace/gohelper/server/grpc/_example/server"
)

func TestHello(t *testing.T) {
	const (
		seconds = 10
	)
	timer := time.NewTimer(time.Second * seconds)
	quit := make(chan struct{})

	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	go server.NewServer()
	go client.NewClient()

	for {
		select {
		case <-quit:
			return
		default:
		}
	}
}
