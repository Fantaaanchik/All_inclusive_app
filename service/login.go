package service

import (
	"all_inclusive_app/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *Service) Login(user models.LoginStruct) (string, error) {
	foundedUser, err := s.Repository.Login(user)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Email: foundedUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		wrappedErr := fmt.Errorf("could not generate token, Error description: %s", err.Error())
		return "", wrappedErr
	}

	return tokenString, nil
}
