package handlers

import (
	"context"
	"net/url"

	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/infra"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/infra/models"
)

func readUserService(ctx context.Context, dto url.Values) (userObj []models.User, err error) {
	var user []models.User

	if err = infra.ReadUser(ctx, &user, dto); err != nil {
		return []models.User{}, nil
	}

	return user, nil
}
