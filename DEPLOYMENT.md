# 🚀 Deployment Guide: Matiks Leaderboard Platform

This guide covers the process of moving your application from local development to a production environment.

---

## 🎨 Part 1: Frontend Deployment (Vercel)

Vercel is the recommended platform for the Expo web frontend.

### 1. Prepare your Expo Web Build
First, ensure your web build is export-ready:
```bash
cd frontend
npx expo export:web
```

### 2. Connect to GitHub
- Log into [Vercel](https://vercel.com).
- Click **"New Project"**.
- Select your repository: `Matiks-Internship-Assignment`.

### 3. Configure Framework & Build
In the Vercel project settings:
- **Framework Preset**: Other (or Expo)
- **Build Command**: `cd frontend && npm install && npx expo export:web`
- **Output Directory**: `frontend/web-build`
- **Root Directory**: `./`

### 4. Environment Variables
Add an Environment Variable named `EXPO_PUBLIC_API_URL` and set it to your deployed backend URL.
> **Note**: If you haven't deployed the backend yet, you can skip this and update it later.

---

## ⚙️ Part 2: Backend Deployment (Railway or Render)

Because the Go backend uses **in-memory storage** (The Frequency Bucket strategy), it requires a "persistent" runtime rather than a "serverless" one (Vercel Functions).

### Recommended: Railway.app
1. Go to [Railway.app](https://railway.app).
2. Create a **New Project** -> **GitHub**.
3. Select your repository.
4. Railway will automatically detect the `backend/main.go` and start the server.
5. In the backend settings, set the **Port** to `8080`.

### Recommended: Render.com
1. Go to [Render.com](https://render.com).
2. Create a **New Web Service**.
3. Point to the `backend/` directory.
4. Build Command: `go build -o main main.go`
5. Start Command: `./main`

---

## 📡 Part 3: Connecting the two

1. Copy the URL of your deployed Backend (e.g., `https://backend-production.up.railway.app`).
2. Go to your **Vercel Frontend Settings**.
3. Update the `EXPO_PUBLIC_API_URL` environment variable.
4. Redeploy the frontend.
1. In the backend code, ensure CORS allows your Vercel domain (the current code uses `*` which is fine for development but should be restricted in production).

---

## 💡 Pro Tip
If you want to keep everything on **Vercel (including Backend)**, you would need to:
1. Move the Go logic to a root `/api` folder.
2. Replace the in-memory frequency buckets with a database (like Redis or PostgreSQL) because Vercel Serverless functions reset their memory every few minutes.

---
**Deployment support for Matiks Assignment**
