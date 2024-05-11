package rest

import (
	"net/http"
	"strings"
	"user_service/app/domain"

	"github.com/gin-gonic/gin"
)

func (r *rest) SignUp(c *gin.Context) {

	var req domain.Users

	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	resp, err := r.Auth.SignUp(c, req)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "Email verification link has been sent", resp, nil)
}

func (r *rest) SignIn(c *gin.Context) {

	var req domain.Auth

	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	resp, err := r.Auth.SignIn(c, req)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "Success Sign In", resp, nil)
}

func (r *rest) SignOut(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	auth := strings.Split(authHeader, " ")
	if len(auth) != 2 || auth[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
		return
	}

	accessToken := auth[1]
	err := r.Auth.SignOut(c, accessToken)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "Success Sign Out", nil, nil)
}

func (r *rest) VerifyEmail(c *gin.Context) {

	tokenHash := c.Query("code")
	verifyType := c.Query("type")

	resp, err := r.Auth.VerifyAuth(c, tokenHash, verifyType)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "Success Verify Email", resp, nil)
}

func (r *rest) ResetPassword(c *gin.Context) {

	var req domain.ResetPassword

	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	err := r.Auth.ResetPassword(c, req.Email)
	if err != nil {
		panic(err)
	}

	r.ResponseJson(c, "Email has been sent", nil, nil)
}

func (r *rest) SuccessAuthenticated(c *gin.Context) {
	r.ResponseJson(c, "Your email already verified, please Sign In", nil, nil)
}
