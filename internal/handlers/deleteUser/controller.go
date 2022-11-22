package handlers

import (
	"context"
	"net/http"

	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/common"
)

type FuncResult struct {
	Error  error
	Result bool
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	defer r.Body.Close()

	ctx, cancel := context.WithTimeout(r.Context(), common.TimeoutQueryDatabase)
	defer cancel()

	resultChannel := make(chan FuncResult)
	defer close(resultChannel)

	go func() {
		ok, err := deleteUserService(ctx, r.Body)
		resultChannel <- FuncResult{Result: ok, Error: err}
	}()

	for {
		select {
		case <-ctx.Done():
			http.Error(w, common.ErrorHandlerTimeout, http.StatusRequestTimeout)
			return
		case result := <-resultChannel:
			if result.Error != nil {
				http.Error(w, common.ErrorBadRequest, http.StatusBadRequest)
				return
			}
			if !result.Result {
				http.Error(w, common.ErrorNotFound, http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}

	}

}
