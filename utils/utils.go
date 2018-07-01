package utils

import (
	"fmt"
)

func BuildURL(ip string, port int) string {
	return fmt.Sprintf("http://%s:%d", ip, port)
}
