package array

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	fmt.Println(Contains([]string{"a", "b", "c"}, "b"))
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]int{1, 2, 3}, 10))
}
