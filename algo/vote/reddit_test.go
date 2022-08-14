package vote_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/algo/vote"
)

func Test_Reddit(t *testing.T) {
	reddit := vote.NewReddit()
	fmt.Println(reddit.Hot(10, 1, time.Now()))
	fmt.Println(reddit.Hot(9, 1, time.Now()))
	fmt.Println(reddit.Hot(8, 1, time.Now()))
	fmt.Println(reddit.Hot(7, 1, time.Now()))
	fmt.Println(reddit.Hot(6, 1, time.Now()))
}
