package server

import (
	"all_inclusive_app/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func Validation(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", token.Header["alg"])
		}
		return []byte(config.Configure.Secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}
