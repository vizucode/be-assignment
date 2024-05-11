package rest

import (
	"fmt"
	"net/http"
	authhandling "user_service/app/middlewares/auth_handling"
	"user_service/app/usecase"

	"github.com/gin-gonic/gin"
)

type rest struct {
	Auth        usecase.IAuth
	Account     usecase.IAccount
	UserAccount usecase.IUserAccount
	Transaction usecase.ITransaction
}

func NewRest(
	Auth usecase.IAuth,
	Account usecase.IAccount,
	UserAccount usecase.IUserAccount,
	Transaction usecase.ITransaction,
) *rest {
	return &rest{
		Auth:        Auth,
		Account:     Account,
		UserAccount: UserAccount,
		Transaction: Transaction,
	}
}

func (r *rest) ResponseJson(ctx *gin.Context, message string, data interface{}, metadata interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  message,
		"data":     data,
		"metadata": metadata,
	})
}

func (r *rest) Register(server *gin.Engine) {
	v1 := server.Group("/api/v1")

	v1.POST("/signup", r.SignUp)
	v1.POST("/signin", r.SignIn)
	v1.POST("/signout", r.SignOut)
	v1.GET("/verify-email", r.VerifyEmail)
	v1.GET("/email-authenticated", r.SuccessAuthenticated)
	v1.POST("/reset-password", r.ResetPassword)

	v1.GET("/types/accounts", authhandling.AuthMiddleware(), r.GetAllTypeAccount)
	v1.GET("/accounts", authhandling.AuthMiddleware(), r.GetAllUserAccount)
	v1.POST("/accounts", authhandling.AuthMiddleware(), r.CreateUserAccount)
	v1.GET("/transactions", authhandling.AuthMiddleware(), r.GetAllTransaction)

	v1.GET("/dummy", authhandling.AuthMiddleware(), func(ctx *gin.Context) {

		email, isExist := ctx.Get("email")
		if isExist {
			fmt.Println(email)
		}

		r.ResponseJson(ctx, "dummy", nil, nil)
	})
}
