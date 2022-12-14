package netx_test

import (
	"fmt"
	"testing"

	"github.com/caicaispace/gohelper/netx"
)

func Test_LocalIP(t *testing.T) {
	netx.LocalIP()
}

func Test_Ip2long(t *testing.T) {
	fmt.Println(netx.Ip2long("192.168.1.1"))
}

func Test_Long2ip(t *testing.T) {
	fmt.Println(netx.Long2ip(3232235777))
}
