package postgresql

import (
	"context"
	"net/http"
	"user_service/app/models"
	customerror "user_service/config/custom_error"
)

func (psql *postgre) UpdateUser(ctx context.Context, selectedField []string, user models.Users) (resp models.Users, err error) {

	if err = psql.db.Model(&models.Users{}).Select(selectedField).Where("email = ?", user.Email).Updates(&user).Error; err != nil {
		return models.Users{}, customerror.NewCustomError(err.Error(), http.StatusInternalServerError)
	}
	return
}
