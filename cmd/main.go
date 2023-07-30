package main

import (
	"fmt"
	"github.com/bednyak/go-react-jwt-auth/pkg/models"
	"github.com/bednyak/go-react-jwt-auth/pkg/routes"
	"github.com/gorilla/handlers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	godotenv.Load(".env")

	models.ConnectToDb()
	routes.InitializeRoute()

	fmt.Printf("Server started at http://localhost:%s", os.Getenv("PORT"))
	err := http.ListenAndServe(
		":"+os.Getenv("PORT"),
		handlers.CORS(
			handlers.AllowedHeaders(
				[]string{
					"X-Requested-With",
					"Access-Control-Allow-Origin",
					"Content-Type", "Authorization",
				}),
			handlers.AllowedMethods(
				[]string{
					"GET",
					"POST",
					"PUT",
					"DELETE",
					"HEAD",
					"OPTIONS",
				}),
			handlers.AllowedOrigins(
				[]string{"*"}))(routes.GetRouter()))
	if err != nil {
		log.Fatal(err)
	}
}
