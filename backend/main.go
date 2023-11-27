package main

import (
	apistream "SLtwitchApi/api/stream"
	"SLtwitchApi/db"
	"SLtwitchApi/models"
	"SLtwitchApi/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}
	migrate := os.Getenv("MIGRATE")
	// Initialize Database
	db.InitDB()
	if migrate == "true" {
		db.SQL.AutoMigrate(&models.Streamer{})
		fmt.Println("Tabelle 'Streamer' erstellt")
	}
	// Initialize CheckStreamer
	services.InitStreamerCall()
	// Initialize MuxRouter
	r := mux.NewRouter()
	// Initialize SubRoutes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/addstream", apistream.AddStreamHandler)
	// api.HandleFunc("/streamers/online", apistream.ListActiveStreams).Methods("GET")
	api.HandleFunc("/streamers/all", apistream.ListAllActiveStreamers).Methods("GET")
	api.HandleFunc("/streamers/details/{username}", GetStreamerDetailsByUsernameHandler).Methods("GET")
	api.Use(loggingMiddleware)
	api.Use(myCorsHandler)

	// Start the Server and Listen!
	fmt.Println("Api-Server started at: http://localhost:7080")
	log.Fatalln(http.ListenAndServe("localhost:7080", r))
}

func GetStreamerDetailsByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the username from the request path
	username := mux.Vars(r)["username"]

	// Call the original function to get the streamer details
	streamer := services.GetStreamerDetailsByUsername(username)

	// Convert the result to JSON and send it in the response
	json.NewEncoder(w).Encode(streamer)
}
