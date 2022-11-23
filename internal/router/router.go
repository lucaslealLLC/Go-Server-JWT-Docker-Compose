package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/common"
	userCreate "github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/handlers/createUser"
	userDelete "github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/handlers/deleteUser"
	getJwt "github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/handlers/getJwt"
	userRead "github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/handlers/readUser"
	userUpdate "github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/handlers/updateUser"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/middlewares"
)

func Router() {

	// TODO: Create tests

	r := mux.NewRouter()

	r.HandleFunc("/users", middlewares.ValidateJwt(userCreate.CreateUserHandler)).Methods(http.MethodPost)
	r.HandleFunc("/users", middlewares.ValidateJwt(userRead.ReadUserHandler)).Methods(http.MethodGet)
	r.HandleFunc("/users", middlewares.ValidateJwt(userDelete.DeleteUserHandler)).Methods(http.MethodDelete)
	r.HandleFunc("/users", middlewares.ValidateJwt(userUpdate.UpdateUserHandler)).Methods(http.MethodPut)

	r.HandleFunc("/jwt", getJwt.GetJwt).Methods(http.MethodGet) // generate jwt cookie

	rWithTimeOut := http.TimeoutHandler(r, common.TimeoutHandler, common.ErrorHandlerTimeout)

	err := http.ListenAndServe(":7000", rWithTimeOut)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Server listening on http://localhost:7000")
	}
}
