package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/common"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/infra"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/infra/models"
)

func updateUserService(ctx context.Context, dto io.ReadCloser) (userObj models.User, err error) {
	var user models.User

	request, err := io.ReadAll(dto)
	if err != nil {
		return models.User{}, err
	}

	if err = json.Unmarshal(request, &user); err != nil {
		return models.User{}, err
	}

	err = validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.By(common.AlphaWithSpaces)),
		validation.Field(&user.Surname, validation.By(common.AlphaWithSpaces)),
	)
	if err != nil {
		return models.User{}, err
	}

	dataUpdated, err := infra.UpdateUser(ctx, &user)

	if err != nil {
		return models.User{}, err
	}

	if !dataUpdated {
		return models.User{}, errors.New(common.ErrorNotFound)
	}

	return user, nil
}
