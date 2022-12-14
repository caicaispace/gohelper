package utils

import "regexp"

// 数字长度
func NumberLen[T int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8](a T) int {
	i := 0
	for a > 0 {
		a /= 10
		i++
	}
	return i
}

// 是否是合法的经纬度
func IsLegalOfLatAndLon(lat, lon string) bool {
	matchLat, _ := regexp.MatchString("^-?((0|1?[0-8]?[0-9]?)(([.][0-9]{1,10})?)|180(([.][0]{1,10})?))$", lat)
	matchLon, _ := regexp.MatchString("^-?((0|[1-8]?[0-9]?)(([.][0-9]{1,10})?)|90(([.][0]{1,10})?))$", lon)
	if matchLat && matchLon {
		return true
	}
	return false
}
