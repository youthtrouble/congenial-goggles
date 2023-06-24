package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/youthtrouble/congenial-goggles/handlers"
	"github.com/youthtrouble/congenial-goggles/telegram"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/oanda", handlers.OandaHandler).Methods("GET")

	// go telegram.InitAlfredTelegramListening()
	telegram.InitOandaTelegramListening()

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

}
