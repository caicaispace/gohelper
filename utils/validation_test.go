package utils

import (
	"fmt"
	"testing"
)

func TestIsLegalOfLatAndLon(t *testing.T) {
	lat := "121.71717171"
	lon := "31.18181818"
	ret := IsLegalOfLatAndLon(lat, lon)
	fmt.Println(ret)
}
