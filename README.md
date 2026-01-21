# Matiks Leaderboard Platform 🏆

A high-performance, aesthetically premium leaderboard system built with **Golang** (Backend) and **React Native Expo** (Frontend). Designed to handle millions of users with **O(1) ranking complexity** and a state-of-the-art user experience.

---

## ✨ Key Highlights

### ⚡ Performance & Scalability
- **Frequency Bucket Algorithm**: Native implementation in Go that allows calculating any user's rank in constant time `O(1)`, regardless of whether there are 10,000 or 10,000,000 users.
- **Efficient Sorting**: The system performs minimal sorting operations, relying on a mathematical bucket strategy to sum frequencies for lightning-fast rank retrieval.
- **Dense Ranking**: Correctly handles ties where users with the same score share the same rank.

### 🎨 Premium User Experience
- **Modern Aesthetic**: Rich glassmorphism effects, smooth gradients, and a curated color palette.
- **Dynamic Dark Mode**: Fully responsive theme that adapts to system preferences or manual toggles.
- **Micro-animations**: Subtle interactions and loading states for a polished "Grandmaster" feel.
- **Search & Simulation**: Real-time player search and a traffic simulation tool to see the leaderboard update live.

---

## 🛠 Tech Stack

| Component | Technology |
| :--- | :--- |
| **Backend** | Golang (Standard Library, Frequency Buckets) |
| **Frontend** | React Native, Expo, TypeScript |
| **Styling** | Vanilla CSS (StyleSheet) for maximum flexibility |
| **Dev Tools** | Metro Bundler, Webpack |

---

## 📂 Project Structure

```text
.
├── backend/
│   ├── core/           # Business logic: Ranking algorithm & Leaderboard state
│   ├── api/            # REST API: HTTP Handlers (Leaderboard, Search, Simulate)
│   └── main.go         # Application entry & Data seeding (10,000 users)
├── frontend/
│   ├── services/       # API layer (fetch, search, simulate)
│   ├── components/     # UI Components
│   └── App.tsx         # Main entry point & Theme management
└── README.md
```

---

## 🚀 Getting Started

### 1. Backend Setup (Go)
The backend seeds **10,000 users** with random ratings (100–5000) on startup.

```bash
cd backend
go run main.go
```
> **Server URL**: `http://localhost:8080`

### 2. Frontend Setup (Expo Web)
```bash
cd frontend
npm install
npx expo start --web
```
> **Default Hub**: `http://localhost:8081` (or next available port)

---

## 📡 API Endpoints

- **GET `/leaderboard?limit=N`**: Fetches the top N players.
- **GET `/search?q=username`**: Searches for a specific user and returns their global rank.
- **GET `/simulate`**: Triggers a random rating update across the dataset to simulate active traffic.

---

## 🧠 Architecture: The Bucket Strategy

To solve the "Millions of Users" requirement, we avoid traditional `O(N log N)` sorting on every request. Instead, we maintain an array where each index represents a possible score (100–5000).

```go
// Rank calculation simplified
func getRank(score int) int {
    rank := 1
    for s := 5000; s > score; s-- {
        rank += bucket[s]
    }
    return rank
}
```
This ensures that fetching the leaderboard or checking a user's rank remains incredibly fast even as the user count scales exponentially.

---
**Developed for Matiks Assignment**
