package utils

import (
	"net"
	"net/http"
	"strings"
)

// Estrae il tipo MIME da una richiesta http (valore di "Accept:")
func RequestMime(header http.Header) string {
	/*if strings.Split(header.Get("Accept"), ",")[0] == "text/html" {
		return "application/json"
	}*/
	return strings.Split(header.Get("Accept"), ",")[0]
}

// Estrae l'indirizzo IP di origine di una http.Request
func RequestIP(r *http.Request) net.IP {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil
	}
	return net.ParseIP(ip)
}

// Trasforma un int64 in stringa
func I64toa(n int64) string {
	buf := [11]byte{}
	pos := len(buf)
	signed := n < 0
	if signed {
		n = -n
	}
	for {
		pos--
		buf[pos], n = '0'+byte(n%10), n/10
		if n == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
