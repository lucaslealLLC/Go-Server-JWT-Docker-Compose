package handlers

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJwt() (jwtToken string, err error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 25).Unix() // 25 minutes to expire

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))

	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}
