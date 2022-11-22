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
	Result []models.User
}

func ReadUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), common.TimeoutQueryDatabase)
	defer cancel()

	resultChannel := make(chan FuncResult)
	defer close(resultChannel)

	go func() {
		user, err := readUserService(ctx, r.URL.Query())
		resultChannel <- FuncResult{Error: err, Result: user}
	}()

	for {
		select {
		case <-ctx.Done():
			http.Error(w, common.ErrorHandlerTimeout, http.StatusRequestTimeout)
			return
		case result := <-resultChannel:
			if result.Error != nil {
				log.Println(result.Error)
				http.Error(w, common.ErrorBadRequest, http.StatusBadRequest)
				return
			}
			if len(result.Result) == 0 {
				http.Error(w, "[]", http.StatusNotFound)
				return
			}
			response, _ := json.Marshal(result.Result)
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}

}
