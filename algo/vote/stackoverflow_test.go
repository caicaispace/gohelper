package vote_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
)

func Test_StackOverflow(t *testing.T) {
	so := vote.NewStackOverflow()
	fmt.Println(so.Hot(0, 0, 3, 3, time.Now(), time.Now()))
	for i := 1; i <= 10; i++ {
		fmt.Println(so.Hot(i, 10, 3, 3, time.Now(), time.Now()))
	}
}
