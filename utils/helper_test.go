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

func TestIsLegalOfLatAndLon(t *testing.T) {
	lat := "121.71717171"
	lon := "31.18181818"
	ret := IsLegalOfLatAndLon(lat, lon)
	fmt.Println(ret)
}
