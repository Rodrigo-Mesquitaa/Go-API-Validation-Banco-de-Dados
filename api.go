package main 

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"api/routes"
	"api/models"
	"api/services/email"
	"github.com/gorilla/handlers"
)

func main() {
	models.TestConnection()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printin("Listening port" + port)
	email.InitializeEmailAuth()
	headers := handlers.AllowedHeaders([]string{"x-Request", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}) 
	origins := handlers.AllowedOrigins([]string{"*"})
  	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headers, methods, origins)(r)))
	}
	