package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unicode/utf8"

	"github.com/TylerBrock/colorjson"
	"github.com/google/uuid"
)

// ServerCloseHandler 服务关闭处理
func ServerCloseHandler(done chan bool, callback func()) {
	fmt.Println("starting ...")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signals)
	go func() {
		<-signals
		done <- true
		fmt.Println("")
		fmt.Println("Ctrl+C pressed in Terminal")
	}()
	<-done
	if callback != nil {
		callback()
	}
	fmt.Println("exiting ...")
}

// NumberLen 数字类型长度
func NumberLen[T int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8](a T) int {
	i := 0
	for a > 0 {
		a /= 10
		i++
	}
	return i
}

// RandInt 生成随机数
func RandInt(min, max int) int {
	// if min >= max || min == 0 || max == 0 { // 防止除0错误
	if min >= max {
		return max
	}
	return min + rand.Intn(max-min)
}

// UnicodeDecode 将 Unicode 字符串解码为普通字符串
func UnicodeDecode(str string) string {
	result := ""
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		result += string(r)
		str = str[size:]
	}
	return result
}

// UnicodeEncode 将普通字符串编码为 Unicode 字符串
func UnicodeEncode(str string) string {
	result := ""
	for _, r := range str {
		result += fmt.Sprintf("\\u%04X", r)
	}
	return result
}

// UnicodeDecode 将 Unicode 字符串解码为普通字符串
func UnicodeDecodeV2(str string) string {
	var builder strings.Builder
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		builder.WriteRune(r)
		str = str[size:]
	}
	return builder.String()
}

// UnicodeEncode 将普通字符串编码为 Unicode 字符串
func UnicodeEncodeV2(str string) string {
	var builder strings.Builder
	for _, r := range str {
		builder.WriteString(fmt.Sprintf("\\u%04X", r))
	}
	return builder.String()
}

// 随机一个 float64
func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// 打印json
func PrintJson(v interface{}) {
	byteData, _ := json.Marshal(v)
	var obj map[string]interface{}
	json.Unmarshal([]byte(byteData), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 4
	ss, _ := f.Marshal(obj)
	fmt.Println(string(ss))
}

// 生成md5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 获取uuid
func Uuid() string {
	return uuid.New().String()
}

// 获取uuid(无 -)
func UuidNoDash() string {
	uuid := uuid.New().String()
	uuid = strings.ReplaceAll(uuid, "-", "")
	return uuid
}

// 获取uuid(8位)
func Uuid8() string {
	uuid := uuid.New().String()
	// uuid = strings.ReplaceAll(uuid, "-", "")
	// fmt.Println("uuid: ", uuid)
	return uuid[:8]
}
