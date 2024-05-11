package main

import (
	"os"
	"strings"
	errorhandling "user_service/app/middlewares/error_handling"
	"user_service/app/repositories/postgresql"
	"user_service/app/repositories/supabase"
	"user_service/app/router/rest"
	account "user_service/app/usecase/account"
	"user_service/app/usecase/auth"
	"user_service/app/usecase/transaction"
	useraccount "user_service/app/usecase/user_account"
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
	supabaseClient := SUPABASE.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API_KEY"), true)

	// repository
	postgresql := postgresql.NewDatabase(dbcon)
	supabase := supabase.NewSupabase(supabaseClient)

	// usecase
	auth := auth.NewAuth(supabase, postgresql)
	userAccount := useraccount.NewUserAccount(postgresql)
	account := account.NewAccount(postgresql)
	transaction := transaction.NewTransaction(postgresql)

	rest.NewRest(
		auth,
		account,
		userAccount,
		transaction,
	).Register(server)

	server.Run(port)
}
