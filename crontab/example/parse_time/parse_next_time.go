package main

import (
	"fmt"
	"time"

	"github.com/caicaispace/gohelper/crontab"
)

func main() {
	t, err := crontab.Parse("0 0 0 */1 * *")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(" now: ", time.Now())
	next := t.Next(time.Now())
	fmt.Println("next: ", next)
}
