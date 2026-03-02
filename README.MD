# Backend Article Portal API

Simple REST API for managing articles built using **Golang**, **MySQL**, and **Clean Architecture (Handler → Service → Repository)**.

---

## 🚀 Tech Stack

- Golang (net/http)
- MySQL
- Viper (Environment Config)
- golang-migrate (Database Migration)

---

## 📦 Features

- Create Article
- Get All Articles (limit & offset)
- Get Article by ID
- Update Article
- Delete Article
- API Key Middleware
- Request Validation

---

## ⚙️ Requirements

Make sure you have installed:

- Go >= 1.21
- MySQL
- golang-migrate CLI

Install migrate:

```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# ============================================================================================

Setup Project (Local Development)
1️⃣ Clone Repository

git clone <your-repository-url>
cd backend-portal
cp ENV_EXAMPLE .env
edit .env

DB_CONN=username:password@/database_name (e.g root:root@/article_post)
PORT=8080
API_KEY=secretweapon


CREATE DATABASE article_post;

migrate -path migrations \
-database "mysql://root:root@tcp(localhost:3306)/article_post" \
up

go run main.go

# ============================================================================================

Postman collection https://tani-permana.postman.co/workspace/HADI-SELF-PROJECT~abc164be-cc16-40d7-86e8-93ce34373872/collection/34741900-a52511ec-a8bb-405e-83a9-9e1cb671d1ca?action=share&source=copy-link&creator=34741900&active-environment=7e74dba0-d467-41c5-b88b-c74b85a00d3a


🧪 API Endpoint
Health Check
GET /api/health

Response:

{
  "status": "OK",
  "message": "API running"
}

Create Article
POST /api/v1/article

Body:

{
  "title": "Article title minimal twenty characters",
  "content": "Content minimal 200 characters...",
  "category": "Tech",
  "status": "publish"
}
Get All Articles
GET /api/v1/article?limit=10&offset=0
Get Article by ID
GET /api/v1/article/{id}
Update Article
PUT /api/v1/article/{id}
Delete Article
DELETE /api/v1/article/{id}

🔐 API Key

Send API Key via header:

X-API-KEY: [your env API_KEY]


deployed project : https://queless.space/dashboard