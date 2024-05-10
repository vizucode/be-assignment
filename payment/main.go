package main

import (
	"os"
	errorhandling "paymentservice/app/middlewares/error_handling"
	"paymentservice/app/router/rest"
	"strings"

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
