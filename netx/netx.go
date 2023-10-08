package netx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// GetClientIP 获取客户端ip
func GetClientIP(r *http.Request) string {
	headers := r.Header
	ip := r.RemoteAddr
	if xRealIP := headers.Get("X-Real-IP"); xRealIP != "" {
		ip = xRealIP
	} else if xForwardedFor := headers.Get("X-Forwarded-For"); xForwardedFor != "" {
		ip = strings.Split(xForwardedFor, ",")[0]
	} else if httpXForwardedFor := headers.Get("HTTP_X_FORWARDED_FOR"); httpXForwardedFor != "" {
		ip = strings.Split(httpXForwardedFor, ",")[0]
	}
	if ip == "" {
		return ""
	}
	// 如果 IP 是一个地址和端口的组合，则只保留 IP 部分
	if index := strings.Index(ip, ":"); index != -1 {
		ip = ip[:index]
	}
	return ip
}

// LocalIP 获取本地ip
func LocalIP() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addr {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// Ip2long 将字符串形式的ip转换成整数形式
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// Long2ip 将整数形式的ip转换成字符串形式
func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

// GetMainFromDomain 获取域名主域名
func GetMainFromDomain(domain string) (string, error) {
	re := regexp.MustCompile(`(?i)[a-z0-9][a-z0-9\-]{0,62}\.[a-z\.]{2,6}$`)
	match := re.FindString(domain)
	if match == "" {
		return "", errors.New("Invalid domain")
	}
	return match, nil
}

// IsLocalIP 判断是否是本地ip
func IsLocalIP(ipAddress string) bool {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return false
	}

	// 判断是否是 IPv4 地址
	if ip.To4() != nil {
		// 判断是否是私有地址段
		privateRanges := []struct {
			start net.IP
			end   net.IP
		}{
			{net.ParseIP("10.0.0.0"), net.ParseIP("10.255.255.255")},
			{net.ParseIP("172.16.0.0"), net.ParseIP("172.31.255.255")},
			{net.ParseIP("192.168.0.0"), net.ParseIP("192.168.255.255")},
		}

		for _, pr := range privateRanges {
			if bytes.Compare(ip, pr.start) >= 0 && bytes.Compare(ip, pr.end) <= 0 {
				return true
			}
		}
	}

	return false
}

// IsIP 判断是否为IP
func IsIP(ip string) bool {
	ipPattern := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	match, _ := regexp.MatchString(ipPattern, ip)
	return match
}
