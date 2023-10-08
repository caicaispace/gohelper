package utils

import "regexp"

// IsLegalOfLatAndLon 判断经纬度是否合法
func IsLegalOfLatAndLon(lat, lon string) bool {
	matchLat, _ := regexp.MatchString("^-?((0|1?[0-8]?[0-9]?)(([.][0-9]{1,10})?)|180(([.][0]{1,10})?))$", lat)
	matchLon, _ := regexp.MatchString("^-?((0|[1-8]?[0-9]?)(([.][0-9]{1,10})?)|90(([.][0]{1,10})?))$", lon)
	if matchLat && matchLon {
		return true
	}
	return false
}
