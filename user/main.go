package main

import (
	"os"
	"strings"
	errorhandling "user_service/app/middlewares/error_handling"
	"user_service/app/router/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	port := os.Getenv("PORT")
	if strings.EqualFold(port, "") {
		port = ":8081"
	}

	server.Use(errorhandling.ErrorHandler())

	rest.NewRest().Register(server)

	server.Run(port)
}
