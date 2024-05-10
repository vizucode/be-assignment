package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type rest struct {
}

func NewRest() *rest {
	return &rest{}
}

func (r *rest) ResponseJson(ctx *gin.Context, message string, data interface{}, metadata interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  message,
		"data":     data,
		"metadata": metadata,
	})
}

func (r *rest) Register(server *gin.Engine) {

}
