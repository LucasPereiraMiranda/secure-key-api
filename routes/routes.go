package routes

import (
	"fmt"
	"log"
	"net/http"
	"password-generator-api/controllers"
	middleware "password-generator-api/middlewares"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router(host string, port string) {
	router := mux.NewRouter()

	router.Use(middleware.ContentTypeMiddleware)

	router.HandleFunc("/api/health", controllers.Health)
	router.HandleFunc("/api/generate-password", controllers.GeneratePassword).Methods("Get")

	address := fmt.Sprintf("%s:%s", host, port)
	log.Fatal(http.ListenAndServe(address, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
