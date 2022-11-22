package infra

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"

	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/infra/models"
	"github.com/lucaslealLLC/Go-Server-JWT-Docker-Compose/internal/initializers"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	DB = initializers.ConnectDatabase()
}

func CreateUser(ctx context.Context, user *models.User) (err error) {

	if err = DB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func ReadUser(ctx context.Context, user *[]models.User, dto url.Values) error {

	parsedQueryParams := make(map[string]string)

	for val := range dto {
		parsedQueryParams[val] = dto[val][0]
	}

	if err := DB.WithContext(ctx).Where(parsedQueryParams).Find(user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, id interface{}) (dataDeleted bool, err error) {

	switch val := id.(type) {
	case string:
		intId, err := strconv.Atoi(val)
		if err != nil {
			return false, errors.New("string is not numerical")
		}
		deleted := DB.WithContext(ctx).Delete(&models.User{}, intId)
		if deleted.Error != nil {
			return false, errors.New("Something went wrong: no data deleted")
		}
		if deleted.RowsAffected == 0 {
			return false, nil
		}
		return true, nil
	case uint, int, float64, float32:
		deleted := DB.WithContext(ctx).Delete(&models.User{}, val)
		if deleted.Error != nil {
			return false, errors.New("Something went wrong: no data deleted")
		}
		if deleted.RowsAffected == 0 {
			return false, nil
		}
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("incorrect data type: %s", reflect.TypeOf(id)))
}

func UpdateUser(ctx context.Context, user *models.User) (dataUpdated bool, err error) {

	updatedResult := DB.WithContext(ctx).Model(user).Where("id = ?", user.ID).Updates(*user)

	if updatedResult.Error != nil {
		return false, updatedResult.Error
	}

	if updatedResult.RowsAffected == 0 {
		return false, nil
	}

	DB.WithContext(ctx).Where("id = ?", user.ID).Find(user)

	return true, nil
}
