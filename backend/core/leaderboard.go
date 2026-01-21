package core

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

// Constants for Rating Limits
const (
	MinRating = 100
	MaxRating = 5000
)

// User represents a player in the leaderboard
type User struct {
	Username string `json:"username"`
	Rating   int    `json:"rating"`
	Rank     int    `json:"rank"` // Dynamically calculated
}

// Leaderboard manages users and rankings efficiently
type Leaderboard struct {
	mu sync.RWMutex

	// Users map: quick lookup by username => O(1)
	Users map[string]*User

	// RatingBuckets: frequency array.
	// buckets[i] stores the count of users who have rating 'i'.
	// Size is MaxRating + 1 to accommodate index 5000 directly.
	RatingBuckets [MaxRating + 1]int
}

// NewLeaderboard initializes the leaderboard
func NewLeaderboard() *Leaderboard {
	return &Leaderboard{
		Users: make(map[string]*User),
	}
}

// AddOrUpdateUser adds a user or updates their score.
// This is O(1) because we just update the map and buckets.
func (l *Leaderboard) AddOrUpdateUser(username string, newRating int) {
	if newRating < MinRating {
		newRating = MinRating
	}
	if newRating > MaxRating {
		newRating = MaxRating
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	user, exists := l.Users[username]
	if exists {
		// Decrement count for old rating
		l.RatingBuckets[user.Rating]--
		// Update user rating
		user.Rating = newRating
	} else {
		// New user
		user = &User{
			Username: username,
			Rating:   newRating,
		}
		l.Users[username] = user
	}

	// Increment count for new rating
	l.RatingBuckets[newRating]++
}

// GetUser returns a user with their LIVE rank.
// Rank calculation is O(1) (technically O(Range), where Range=5000 constant).
func (l *Leaderboard) GetUser(username string) (*User, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	user, exists := l.Users[username]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	// Calculate rank on the fly
	rank := l.calculateRank(user.Rating)

	// Return a copy or updated struct to avoid race conditions if caller modifies it (though we return pointer here, be careful)
	// For API response, we package it nicely.
	return &User{
		Username: user.Username,
		Rating:   user.Rating,
		Rank:     rank,
	}, nil
}

// SearchUsers performs a prefix search on usernames.
// For 10k users, a linear scan is fine (~ms). For millions, we would use a Trie.
// Returns top 10 matches.
func (l *Leaderboard) SearchUsers(query string) []*User {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var results []*User
	query = strings.ToLower(query)

	count := 0
	for _, user := range l.Users {
		if strings.Contains(strings.ToLower(user.Username), query) {
			rank := l.calculateRank(user.Rating)
			results = append(results, &User{
				Username: user.Username,
				Rating:   user.Rating,
				Rank:     rank,
			})
			count++
			if count >= 20 { // Limit results for performance
				break
			}
		}
	}
	// Note: The results from map iteration are random order.
	// In a real system, we might want to sort these results by Rank or Match Quality.
	// For this assignment, returning matches is key.

	// Sort by Rank (Ascending) then Username
	sort.Slice(results, func(i, j int) bool {
		if results[i].Rank == results[j].Rank {
			return results[i].Username < results[j].Username
		}
		return results[i].Rank < results[j].Rank
	})

	return results
}

// GetTopUsers returns the top N users (e.g., for the leaderboard screen).
func (l *Leaderboard) GetTopUsers(limit int) []*User {
	l.mu.RLock()
	defer l.mu.RUnlock()

	// Temporary: Collect all and Sort.
	// For 10k users, this is performant enough.
	all := make([]*User, 0, len(l.Users))
	for _, u := range l.Users {
		all = append(all, &User{Username: u.Username, Rating: u.Rating})
	}

	// Sort by Rating Descending
	sort.Slice(all, func(i, j int) bool {
		if all[i].Rating == all[j].Rating {
			return all[i].Username < all[j].Username
		}
		return all[i].Rating > all[j].Rating
	})

	// Assign Ranks
	for _, u := range all {
		u.Rank = l.calculateRank(u.Rating)
	}

	// Apply Limit
	if limit > len(all) {
		limit = len(all)
	}

	return all[:limit]
}

// calculateRank computes rank from score using buckets. O(1).
// Rank = 1 + (Count of users strictly better than score)
func (l *Leaderboard) calculateRank(rating int) int {
	rank := 1
	// Sum all buckets higher than this rating
	for r := rating + 1; r <= MaxRating; r++ {
		rank += l.RatingBuckets[r]
	}
	return rank
}

// Seed populates the leaderboard with N random users
func (l *Leaderboard) Seed(n int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		username := fmt.Sprintf("user_%d", i+1)
		// Rating 100-5000
		rating := rand.Intn(MaxRating-MinRating+1) + MinRating
		l.AddOrUpdateUser(username, rating)
	}
}
