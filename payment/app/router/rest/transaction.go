package rest

import (
	"net/http"
	"paymentservice/app/domain"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetAllTransaction(c *gin.Context) {

	email, isExist := c.Get("email")
	if !isExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "email not found"})
	}

	transactionResp, err := r.Transaction.GetAllTransaction(c, email.(string))
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "success get all transaction", transactionResp, nil)
}

func (r *rest) Send(c *gin.Context) {
	email, isExist := c.Get("email")
	if !isExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "email not found"})
	}

	var send domain.Send
	err := c.ShouldBindBodyWithJSON(&send)
	if err != nil {
		panic(err)
	}

	send.OwnerEmail = email.(string)

	transactionResp, err := r.Transaction.Send(c, send)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "success send transaction", transactionResp, nil)
}
