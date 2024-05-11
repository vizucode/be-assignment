package postgresql

import (
	"context"
	"errors"
	"net/http"
	"paymentservice/app/models"
	customerror "paymentservice/config/custom_error"

	"gorm.io/gorm"
)

func (psql *postgre) GetAllTypeAccount(ctx context.Context) (resp []models.Account, err error) {

	if err = psql.db.Model(&models.Account{}).Where("status = ?", true).Find(&resp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.Account{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return resp, nil
}

func (psql *postgre) FindTypeAccount(ctx context.Context, typeAccountId int) (resp models.Account, err error) {
	if err = psql.db.Model(&models.Account{}).Where("id = ?", typeAccountId).First(&resp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Account{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return
}

func (psql *postgre) GetUserAccount(ctx context.Context, userId int) (resp []models.UserAccount, err error) {

	if err = psql.db.Model(&models.UserAccount{}).Where("user_id = ?", userId).Preload("Transactions.AccountOwner.User").Preload("Transactions.AccountTarget.User").Find(&resp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return
}

func (psql *postgre) CreateUserAccount(ctx context.Context, userAccount models.UserAccount) (resp models.UserAccount, err error) {

	if err = psql.db.Model(&models.UserAccount{}).Create(&userAccount).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return
}

func (psql *postgre) FindUserAccount(ctx context.Context, userAccountCode string) (resp models.UserAccount, err error) {

	if err = psql.db.Model(&models.UserAccount{}).Where("account_code = ?", userAccountCode).Preload("User").First(&resp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return
}

func (psql *postgre) CheckOwnerAccount(ctx context.Context, userId int, accountCode string) (resp models.UserAccount, err error) {

	if err = psql.db.Model(&models.UserAccount{}).Where("user_id = ?", userId).Where("account_code = ?", accountCode).Preload("User").First(&resp).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserAccount{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
		}
	}

	return resp, nil
}
