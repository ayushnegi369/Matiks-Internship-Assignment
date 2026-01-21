package api

import (
	"encoding/json"
	"math/rand"
	"matiks-leaderboard/core"
	"net/http"
	"strconv"
)

var Board *core.Leaderboard

func HandleLeaderboard(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCors(&w)

	limitStr := r.URL.Query().Get("limit")
	limit := 50 // Default
	if limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil {
			limit = val
		}
	}
	// Cap limit for safety
	if limit > 500 {
		limit = 500
	}

	users := Board.GetTopUsers(limit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	results := Board.SearchUsers(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func HandleSimulate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Randomly update 100 users
	// In a real scenario, this might be triggered by a game engine.
	// We'll update random existing users to simulate activity.

	// Create a goroutine to not block the response?
	// Or just do it synchronously for the "Trigger" effect.

	// Let's do a loop of random updates
	// Note: We need access to keys to pick random ones, but map iteration is random order essentially
	// To be truly random without iterating the whole map, we'd need a keys slice.
	// For simulation, we can just "Iterate slightly" or pick random IDs if we knew the format "user_X"

	count := 0
	Board.Seed(0) // Hack? No.

	// Let's pick random ID user_{rand}
	// 50 updates
	for i := 0; i < 50; i++ {
		// Assuming we have user_1 to user_10000
		id := rand.Intn(10000) + 1
		username := "user_" + strconv.Itoa(id)

		// New Random Rating
		newRating := rand.Intn(core.MaxRating-core.MinRating+1) + core.MinRating

		Board.AddOrUpdateUser(username, newRating)
		count++
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Simulated 50 random updates"})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
}
