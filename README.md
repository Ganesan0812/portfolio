# Portfolio: Vue 3 + Go

portfolio/
  backend/   Go REST API (standard library only)
  frontend/  Vue 3 + Vite UI

## Prerequisites
- Go 1.22+      (go version)
- Node 18+      (node -v)

## Run the backend (terminal 1)
cd backend
go run .
# API on http://localhost:8080
# GET /api/profile | GET /api/projects | POST /api/contact

## Run the frontend (terminal 2)
cd frontend
npm install
npm run dev
# UI on http://localhost:5173  (Vite proxies /api to :8080)

## Production build
cd frontend && npm run build      # outputs frontend/dist
cd ../backend && go build -o portfolio-api .

## Test the API quickly
curl http://localhost:8080/api/projects
curl -X POST http://localhost:8080/api/contact \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"a@b.com","message":"Hello"}'

## Customize
- Edit your details/projects in backend/main.go (profile and projects variables)
- Edit look and layout in frontend/src/App.vue and style.css
