package server

import (
	"all_inclusive_app/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(config.Configure.TokenLifeSeconds)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Configure.Secret))
}
