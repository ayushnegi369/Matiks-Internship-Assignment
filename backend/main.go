package main

import (
	"fmt"
	"log"
	"matiks-leaderboard/api"
	"matiks-leaderboard/core"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 0. Load environment variables from root .env if it exists
	// We look one level up since we run from /backend
	_ = godotenv.Load("../.env")

	// 1. Initialize Leaderboard
	api.Board = core.NewLeaderboard()

	// 2. Seed Data
	fmt.Println("Seeding 10,000 users...")
	api.Board.Seed(10000)
	fmt.Println("Seeding complete.")

	// 3. Define Routes
	http.HandleFunc("/leaderboard", api.HandleLeaderboard)
	http.HandleFunc("/search", api.HandleSearch)
	http.HandleFunc("/simulate", api.HandleSimulate)

	// 4. Start Server
	rawPort := os.Getenv("PORT")
	if rawPort == "" {
		rawPort = "8080"
	}
	port := ":" + rawPort
	fmt.Println("Server starting on port " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
