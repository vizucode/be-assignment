package postgresql

import (
	"context"
	"errors"
	"net/http"
	"paymentservice/app/models"
	customerror "paymentservice/config/custom_error"

	"gorm.io/gorm"
)

func (psql *postgre) GetAllTransaction(ctx context.Context, userId int) (resp []models.Transaction, err error) {

	if err = psql.db.Model(&models.Transaction{}).Where("user_id = ?", userId).Preload("AccountOwner.User").Preload("AccountTarget.User").Find(&resp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.Transaction{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return
}

func (psql *postgre) CreateTransaction(ctx context.Context, transaction models.Transaction) (resp models.Transaction, err error) {

	if err = psql.db.Model(&models.Transaction{}).Create(&transaction).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Transaction{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return
}
