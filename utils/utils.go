package utils

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

const (
	urlRegex = `(https?)://([^:^/]*):?(\d*)?(.*)?`
)

func BuildURL(proto, url string, port int) string {
	if url == "" {
		errors.New("No URL provided.")
	}

	if proto == "" {
		proto = "http"
	}

	if proto == "https" && port == 0 {
		port = 443
	}

	return fmt.Sprintf("%s://%s:%d", proto, url, port)
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
