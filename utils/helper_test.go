package utils

import (
	"fmt"
	"testing"
)

func TestNumberLen(t *testing.T) {
	num := 9999999999999
	fmt.Println(NumberLen(num))
	fmt.Println(NumberLen(int64(num)))
	fmt.Println(NumberLen(int32(num)))
	fmt.Println(NumberLen(int16(num)))
	fmt.Println(NumberLen(int8(num)))
	fmt.Println(NumberLen(int64(num)))
	fmt.Println(NumberLen(uint(num)))
	fmt.Println(NumberLen(uint64(num)))
	fmt.Println(NumberLen(uint32(num)))
	fmt.Println(NumberLen(uint16(num)))
	fmt.Println(NumberLen(uint8(num)))
}
