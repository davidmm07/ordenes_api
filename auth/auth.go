package auth

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql/gqlerrors"
)

var jwtSecret []byte = []byte("bhu2eE31njusecr2etnjimko")

func CrearToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", gqlerrors.FormatError(err)
	}

	return tokenString, nil
}

func ValidarToken(t string) (bool, error) {
	if t == "" {
		return false, gqlerrors.FormatError(errors.New("Token de autorización debe estar presente"))
	}

	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Hubo un error")
		}
		return jwtSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, gqlerrors.FormatError(errors.New("Token de autorización invalida"))
	}
}
