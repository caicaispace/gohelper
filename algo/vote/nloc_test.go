package vote_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
)

func Test_NLOC(t *testing.T) {
	timer := time.NewTimer(time.Second * 10)
	quit := make(chan struct{})

	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	timex := time.Now()
	go func() {
		n := vote.NewNLOC()
		ticker := time.NewTicker(time.Second * 1)
		for {
			<-ticker.C
			fmt.Println(n.Hot(timex))
		}
	}()

	for {
		<-quit
		return
	}
}
