package main

import (
	"fmt"
	"log"
	"matiks-leaderboard/api"
	"matiks-leaderboard/core"
	"net/http"
)

func main() {
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
	port := ":8080"
	fmt.Println("Server starting on port " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
