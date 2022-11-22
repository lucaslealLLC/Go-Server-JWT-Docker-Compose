package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/common"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/infra/models"
)

type FuncResult struct {
	Error  error
	Result models.User
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	defer r.Body.Close()

	ctx, cancel := context.WithTimeout(r.Context(), common.TimeoutQueryDatabase)
	defer cancel()

	resultChannel := make(chan FuncResult)
	defer close(resultChannel)

	go func() {
		user, err := updateUserService(ctx, r.Body)
		resultChannel <- FuncResult{Result: user, Error: err}
	}()

	for {
		select {
		case <-ctx.Done():
			http.Error(w, common.ErrorHandlerTimeout, http.StatusRequestTimeout)
			return
		case result := <-resultChannel:
			if result.Error != nil {
				log.Println(result.Error)
				if result.Error.Error() == common.ErrorNotFound {
					log.Println(result.Error)
					http.Error(w, common.ErrorNotFound, http.StatusNotFound)
					return
				}
				http.Error(w, common.ErrorBadRequest, http.StatusBadRequest)
				return
			}
			response, _ := json.Marshal(result.Result)
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}
}
