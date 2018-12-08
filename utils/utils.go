package utils

import (
	"fmt"
	"os"
	"regexp"
)

const (
	urlRegex = `(https?)://([^:^/]*):?(\d*)?(.*)?`
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
