package vote_test

import (
	"fmt"
	"testing"

	"github.com/caicaispace/gohelper/algo/vote"
)

func Test_WilsonScoreInterval(t *testing.T) {
	wsi := vote.NewWilsonScoreInterval()
	fmt.Println(wsi.Hot(0, 0))
	for i := 1; i <= 10; i++ {
		fmt.Println(wsi.Hot(i, 10))
	}
}
