package main

import (
	"os"
	errorhandling "paymentservice/app/middlewares/error_handling"
	"paymentservice/app/repositories/postgresql"
	"paymentservice/app/repositories/supabase"
	"paymentservice/app/router/rest"
	account "paymentservice/app/usecase/account"
	"paymentservice/app/usecase/auth"
	"paymentservice/app/usecase/transaction"
	useraccount "paymentservice/app/usecase/user_account"
	"paymentservice/config/database"
	"strings"

	SUPABASE "github.com/nedpals/supabase-go"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	port := os.Getenv("PORT")
	if strings.EqualFold(port, "") {
		port = ":8082"
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
