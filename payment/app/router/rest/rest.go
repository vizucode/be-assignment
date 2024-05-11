package rest

import (
	"net/http"
	authhandling "paymentservice/app/middlewares/auth_handling"
	"paymentservice/app/usecase"

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
	// v1 := server.Group("/api/v1")
	v1 := server.Group("/payment/api/v1")

	v1.POST("/send", authhandling.AuthMiddleware(), r.Send)
}
