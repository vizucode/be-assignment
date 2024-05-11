package postgresql

import (
	"context"
	"errors"
	"net/http"
	"paymentservice/app/models"
	customerror "paymentservice/config/custom_error"

	"gorm.io/gorm"
)

func (psql *postgre) CreateUser(ctx context.Context, user models.Users) (resp models.Users, err error) {

	if err = psql.db.Model(&models.Users{}).Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return models.Users{}, customerror.NewCustomError("User already exist", http.StatusInternalServerError)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Users{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return user, nil
}

func (psql *postgre) GetUser(ctx context.Context, email string) (resp models.Users, err error) {
	if err = psql.db.Model(&models.Users{}).Where("email = ?", email).Where("status = ?", true).First(&resp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Users{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return
}

func (psql *postgre) UpdateUser(ctx context.Context, selectedField []string, user models.Users) (resp models.Users, err error) {

	if err = psql.db.Model(&models.Users{}).Select(selectedField).Where("email = ?", user.Email).Updates(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Users{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}
	return
}
