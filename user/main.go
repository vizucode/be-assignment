package main

import (
	"os"
	"strings"
	errorhandling "user_service/app/middlewares/error_handling"
	"user_service/app/repositories/postgresql"
	"user_service/app/repositories/supabase"
	"user_service/app/router/rest"
	"user_service/app/usecase/auth"
	"user_service/config/database"

	SUPABASE "github.com/nedpals/supabase-go"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	port := os.Getenv("PORTS")
	if strings.EqualFold(port, "") {
		port = ":80"
	}
	// configuration
	server.Use(errorhandling.ErrorHandler())
	dbcon := database.GetDB()
	// supabaseClient := SUPABASE.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API_KEY"), true)
	supabaseClient := SUPABASE.CreateClient("https://qytwpijdbksagnkuqgay.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InF5dHdwaWpkYmtzYWdua3VxZ2F5Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTUzNjE4NDIsImV4cCI6MjAzMDkzNzg0Mn0.vml4ckytjDbcsh07VgYSodEJLPVylqU8JGoeA4-X42U", true)

	// repository
	postgresql := postgresql.NewDatabase(dbcon)
	supabase := supabase.NewSupabase(supabaseClient)

	// usecase
	auth := auth.NewAuth(supabase, postgresql)

	rest.NewRest(
		auth,
	).Register(server)

	server.Run(port)
}
