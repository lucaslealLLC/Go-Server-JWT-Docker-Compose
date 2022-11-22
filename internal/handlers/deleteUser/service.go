package handlers

import (
	"context"
	"encoding/json"
	"io"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/common"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/infra"
)

type DeleteRequest struct {
	ID interface{} `json:"id"`
}

func deleteUserService(ctx context.Context, dto io.ReadCloser) (ok bool, err error) {
	var request DeleteRequest

	parsedDto, err := io.ReadAll(dto)
	if err != nil {
		return false, err
	}

	if err = json.Unmarshal(parsedDto, &request); err != nil {
		return false, err
	}

	err = validation.ValidateStruct(&request,
		validation.Field(&request.ID, validation.Required, validation.By(common.IdNumeric)),
	)
	if err != nil {
		return false, err
	}

	dataDeleted, err := infra.DeleteUser(ctx, request.ID)

	if err != nil {
		return false, err
	}

	if !dataDeleted {
		return false, nil
	}

	return true, nil
}
