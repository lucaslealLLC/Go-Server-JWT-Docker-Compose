package middlewares

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/common"
)

func ValidateJwt(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestToken, err := r.Cookie("Token")

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(common.ErrorUnauthorized))
			return
		}

		token, err := jwt.Parse(requestToken.Value, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("incorrect method")
			}

			return []byte(os.Getenv("JWTSECRET")), nil
		})

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(common.ErrorUnauthorized))
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(common.ErrorUnauthorized))
	})
}
