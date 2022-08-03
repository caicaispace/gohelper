package _example_test

import (
	"testing"
	"time"

	"github.com/caicaispace/gohelper/server/grpc/_example"
)

func TestHello(t *testing.T) {
	const (
		seconds = 5
	)
	timer := time.NewTimer(time.Second * seconds)
	quit := make(chan struct{})

	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	go _example.NewServer()
	go _example.NewClient()

	for {
		select {
		case <-quit:
			return
		default:
		}
	}
}
