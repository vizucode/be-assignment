package rest

import "github.com/gin-gonic/gin"

func (r *rest) GetAllTypeAccount(ctx *gin.Context) {

	result, err := r.Account.GetAllTypeAccount(ctx)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(ctx, "Get All Type Account", result, nil)

}
