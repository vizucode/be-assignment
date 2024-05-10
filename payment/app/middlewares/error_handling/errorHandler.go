package errorhandling

import (
	"net/http"
	customerror "paymentservice/config/custom_error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				// reflect err to custom error
				if e, ok := err.(customerror.CustomError); ok {
					ctx.AbortWithStatusJSON(e.HttpCode, gin.H{
						"message": e.Message,
					})
				}

				if validationErr, ok := err.(validator.ValidationErrors); ok {
					// Create a map to hold the validation errors
					errsMap := make(map[string]interface{})
					// Populate the map with validation errors
					for _, val := range validationErr {
						errsMap[val.Field()] = "Terdapat kesalaha " + val.Tag() + " di " + val.Field()
					}

					// Return a response with bad request status and the validation error map
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message": errsMap,
					})
				}
			}
		}()
		ctx.Next()
	}
}
