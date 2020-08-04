/*
 * auth_test.go
 *
 * File di test per il package auth.
 *
 * Copyright (c) 2020 Antonio Napolitano <nap@antonionapolitano.eu>
 *
 * This file is part of webapi-dav
 *
 * webapi-dav is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * webapi-dav is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with webapi-dav. If not, see <http://www.gnu.org/licenses/>.
 */

package auth

import (
	"testing"
	"fmt"
	"flag"

	"github.com/Baldomo/webapi-dav/pkg/config"
)

func TestAll(t *testing.T) {
	// Carica e visualizza la configurazione di test
	configPtr := flag.String("config", "../../config.toml", "Indirizzo del file di configurazione, in .toml o .json")
	err := config.LoadPrefs(*configPtr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config.GetConfig())
	}
	err = InitializeSigning()
	if err != nil {
		fmt.Println(err)
	}

	// Verifica un token
	userInfo, err := ParseToken([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7ImlzcyI6ImRhdiIsInN1YiI6Im5vbWUuY29nbm9tZSIsImF1ZCI6WyJodHRwOi8vZXhhbXBsZS5vcmciLCJodHRwczovL2V4YW1wbGUub3JnIl0sImV4cCI6MjU5MzU0MDQwOSwiaWF0IjoxNTkwOTQ4NDA5fX0.0NARK5Q7Sm_KH_qYPXbu26ihGzYheJHp7-cokhusCEM"))

	if err == nil {
		fmt.Println(userInfo)
	} else {
		fmt.Println(err)
	}
}
