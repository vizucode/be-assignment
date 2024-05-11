package postgresql

import (
	"context"
	"errors"
	"net/http"
	"user_service/app/models"
	customerror "user_service/config/custom_error"

	"gorm.io/gorm"
)

func (psql *postgre) CreateUser(ctx context.Context, user models.Users) (resp models.Users, err error) {

	if err = psql.db.Model(&models.Users{}).Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return models.Users{}, customerror.NewCustomError("User already exist", http.StatusInternalServerError)
		}
		return models.Users{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return user, nil
}
