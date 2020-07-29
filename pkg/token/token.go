package token

import (
	"fmt"
	"math/rand"
	"time"
)

var token = ""

func Get() string {
	return token
}

func Rotate() {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(time.Hour)

	go func() {
		defer ticker.Stop()
		for {
			<-ticker.C
			refreshToken()
		}
	}()
}

func refreshToken() {
	key := make([]byte, 8)
	rand.Read(key)
	token = fmt.Sprintf("%x", key[:8])
}
