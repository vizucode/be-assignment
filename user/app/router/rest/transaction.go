package rest

import (
	"net/http"

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
