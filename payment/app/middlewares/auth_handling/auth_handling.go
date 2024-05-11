package authhandling

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"paymentservice/app/repositories/supabase"

	SUPABASE "github.com/nedpals/supabase-go"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		refreshToken := c.GetHeader("Refresh-Token")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SUPABASE_JWT_SECRET")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				c.Abort()
				return
			}

			if errors.Is(err, jwt.ErrTokenExpired) {
				fmt.Println(err)
				supabaseClient := SUPABASE.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API_KEY"), true)
				authClient := supabase.NewSupabase(supabaseClient)

				result, err := authClient.RefreshToken(context.Background(), tokenString, refreshToken)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"message": "Refresh Token Invalid, Please Sign Up",
					})
					return
				}

				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"data":    result,
					"message": "Token refreshed successfully",
				})
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("email", token.Claims.(jwt.MapClaims)["email"])

		c.Next()
	}
}
