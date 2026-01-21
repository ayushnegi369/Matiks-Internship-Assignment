## 🌍 Centralized Configuration

You can now manage all public URLs and ports from the single **`.env`** file in the project root.

### Environment Variables
- `PORT`: Local port for the Go server (default: 8080).
- `FRONTEND_URL`: Used by the backend to allow CORS from your frontend domain.
- `EXPO_PUBLIC_API_URL`: Used by the Expo app to connect to the backend.

> [!TIP]
> **For Local Dev**: The backend automatically looks for the root `.env` file. For the frontend, I've added a copy of this file to `frontend/.env` so Expo can pick it up.

---

## 🎨 Part 1: Frontend Deployment (Vercel)
...
### 4. Environment Variables
Add your `EXPO_PUBLIC_API_URL` directly in the Vercel Dashboard.

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
