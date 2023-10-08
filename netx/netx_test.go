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

func Test_GetMainDomain(t *testing.T) {
	domains := []string{
		"www/wwwrot/ozbb.cn.log",
		"adnews.cooco.net.cn",
		"http://esphp.chazidian.com",
		"www.xitieba.net",
		"jgcyp.100ky.cn",
		"jfcpt.100ky.cn",
		"diwumeiwen.com",
		"m.zuowen.chazidian.com",
		"scxichen.com",
		"www.jpxue.com",
		"fwzjw.com",
		"jjdije.100ky.cn",
		"static202106.cooco.net.cn",
		"img.easyicon.net",
	}
	for _, domain := range domains {
		mainDomain, err := netx.GetMainFromDomain(domain)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Main domain:", mainDomain)
	}
}
