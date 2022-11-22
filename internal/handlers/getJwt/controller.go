package handlers

import (
	"net/http"
	"os"

	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/common"
)

func GetJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Auth"] != nil {

		if r.Header["Auth"][0] != os.Getenv("JWTSECRET") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(common.ErrorInvalidAuth))
			return
		}

		token, err := CreateJwt()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(common.ErrorSettingToken))
			return
		}

		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			Name:     "Token",
			Value:    token,
		})

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Token has been generated"}`))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(common.ErrorMissingAuth))
}
