package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	router := mux.NewRouter()
	routes := &Routes{}

	router.Use(authMiddleware)
	router.HandleFunc("/addJob", routes.addJob).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("🌐 Server Live at %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func LoadEnv() {
	if os.Getenv("APP_MODE") != "prod" {
		err := godotenv.Load("../../.env") // Loading env variables
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
