package rest

import (
	"net/http"
	"user_service/app/domain"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetAllUserAccount(ctx *gin.Context) {

	email, isExist := ctx.Get("email")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "email not found"})
	}

	result, err := r.UserAccount.GetUserAccount(ctx, email.(string))
	if err != nil {
		panic(err)
	}

	r.ResponseJson(ctx, "Get All User Account", result, nil)
}

func (r *rest) CreateUserAccount(ctx *gin.Context) {

	email, isExist := ctx.Get("email")
	if !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email not found"})
	}

	var userAccount domain.UserAccount
	err := ctx.ShouldBindJSON(&userAccount)
	if err != nil {
		panic(err)
	}

	_, err = r.UserAccount.CreateUserAccount(ctx, email.(string), userAccount)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(ctx, "Create User Account Success", nil, nil)
}
