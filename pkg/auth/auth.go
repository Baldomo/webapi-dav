/*
 * auth.go
 *
 * Funzioni per la verifica dei token JWT di autorizzazione.
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
	"errors"
	"fmt"
	"time"

	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/gbrlsnchs/jwt/v3"
)

var (
	jwtSigner *jwt.HMACSHA
)

// Errore nella verifica del token
type InvalidTokenError struct {
	token []byte
}

func (e *InvalidTokenError) Error() string {
	return fmt.Sprintf("Failed to verify the following token: %s", string(e.token))
}

// Formato del payload JWT
type customPayload struct {
	Payload jwt.Payload
}

// Valore di ritorno di ParseToken
type UserInfo struct {
	Username string `json:"username"`
}

// Inizializza l'algoritmo per la firma HS256
func InitializeSigning() error {

	// Ottiene la chiave segreta dalla configurazione
	secret := config.GetConfig().Auth.JWTSecret

	// Ritorna un errore se manca la chiave per la firma
	if secret == "" {
		return errors.New("Chiave per la firma non specificata!")
	}

	// Inizializza l'algoritmo
	jwtSigner = jwt.NewHS256([]byte(secret))

	return nil
}

// Verifica un token e ne restituisce le informazioni
func ParseToken(token []byte) (UserInfo, error) {

	var (
		// Ottengo il tempo corrente
		now = time.Now()

		// Carico l'FQDN dalla configurazione e definisco l'audience
		fqdn = config.GetConfig().General.FQDN
		aud  = jwt.Audience{"http://" + fqdn, "https://" + fqdn}

		// Inizializzo i "validatori"
		iatValidator = jwt.IssuedAtValidator(now)
		expValidator = jwt.ExpirationTimeValidator(now)
		audValidator = jwt.AudienceValidator(aud)

		// Costruisco il validatore supremo
		pl              customPayload
		validatePayload = jwt.ValidatePayload(&pl.Payload, iatValidator, expValidator, audValidator)
	)

	// Verifico il token
	_, err := jwt.Verify(token, jwtSigner, &pl, validatePayload)

	if err != nil {
		// Valori di errore
		return UserInfo{"1337.h4x0r"}, &InvalidTokenError{token}
	}

	// Ottengo le informazioni sull'utente
	username := pl.Payload.Subject

	return UserInfo{username}, nil
}
