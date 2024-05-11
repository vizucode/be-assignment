package errorhandling

import (
	"net/http"
	customerror "user_service/config/custom_error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				// reflect err to custom error
				if e, ok := err.(*customerror.CustomError); ok {
					ctx.AbortWithStatusJSON(e.HttpCode, gin.H{
						"message": e.Message,
					})
				}

				if validationErr, ok := err.(validator.ValidationErrors); ok {
					// Create a map to hold the validation errors
					errsMap := make(map[string]interface{})
					// Populate the map with validation errors
					for _, val := range validationErr {
						switch val.Tag() {
						case "required":
							errsMap[val.Field()] = "This field is required"
						case "min":
							errsMap[val.Field()] = "This field must be at least 8 characters long"
						case "email":
							errsMap[val.Field()] = "This field must be a valid email address"
						default:
							errsMap[val.Field()] = "This field is invalid"
						}
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
