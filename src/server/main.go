package main

import (
	"net/http"

	"log"

	"server/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello private\n"))
}

func main() {
	loadEnv()
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHandlars := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	r.HandleFunc("/private", middleware.AuthMiddleware(private))

	log.Fatal(http.ListenAndServe(":8888", handlers.CORS(allowedOrigins, allowedMethods, allowedHandlars)(r)))
}
