package utils

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func RequestMime(header http.Header) string {
	/*if strings.Split(header.Get("Accept"), ",")[0] == "text/html" {
		return "application/json"
	}*/
	return strings.Split(header.Get("Accept"), ",")[0]
}

func RequestIP(r *http.Request) net.IP {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil
	}
	return net.ParseIP(ip)
}

func Error(origin string, format string, args ...interface{}) error {
	return fmt.Errorf(origin+format, args...)
}
