package utils

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

const (
	urlRegex = `(https?)://([^:^/]*):?(\d*)?(.*)?`
	Iso86001 = "2006-01-02T15:04:05.000Z"
)

func BuildURL(proto, ip string, port int) string {
	return fmt.Sprintf("%s://%s:%d", proto, ip, port)
}

func SplitUrl(url string) []string {
	pattern := regexp.MustCompile(urlRegex)
	match := pattern.FindStringSubmatch(url)
	return match
}

func IsExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func ConvertISO8601String(s string) string {
	t, err := time.Parse(s, "2006-01-02")
	if err != nil {
		return ""
	}
	return t.Format(Iso86001)
}

func ConvertISO8601(year, month, day, hour, min, sec, nsec int) string {
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, time.UTC).Format(Iso86001)
}
