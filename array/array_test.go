package array_test

import (
	"testing"

	"github.com/caicaispace/gohelper/array"
)

func TestGetLastItem(t *testing.T) {
	slice := []int{1, 2, 3}
	t.Log(array.GetLastItem(slice))
}
