# Auth-API  
A clean-architecture **User Authentication API** built with **Go (Golang)**, **Gin**, **MySQL**, **Redis**, and **JWT**.  
Register, login, refresh token, logout — all secured with JWT & black-list on Redis.

---

## 🚀 Features
- RESTful endpoints (`/api/v1/register`, `/login`, `/refresh`, `/logout`, `/profile`)
- JWT-based stateless auth
- Token black-listing on Redis (logout)
- Auto-migrate database schema
- Clean architecture (domain-driven, layered)
- MySQL & Redis ready
- Docker-friendly

---

## 📦 Tech Stack
| Layer          | Tech                         |
|----------------|------------------------------|
| HTTP Framework | Gin                          |
| ORM            | GORM                         |
| DB             | MySQL 8+                     |
| Cache          | Redis                        |
| Auth           | JWT (HS256)                  |
| Environment    | `.env`                       |
| License        | MIT                          |

---

## 🛠️ Quick Start

### 1. Clone & Install
```bash
git clone https://github.com/yourname/auth-api.git
cd auth-api
go mod tidy
```

### 2. Run App
```bash
go run cmd/main.go
# Server listening on :8080
```

---

### Projek Struktur
```bash
auth-api/
├── cmd/main.go    # entry point
├── internal/
│   ├── config/           
│   ├── domain/           
│   ├── repository/       
│   ├── usecase/          
│   ├── delivery/http/    
│   └── middleware/        
├── pkg/
│   ├── jwt/               
│   └── redisclient/       
├── .env                   
├── go.mod
├── docker-compose.yml
└── README.md
```

### Contributing
We love PRs!
1. Fork → create feature branch → push → open PR.
2. Follow conventional commits & add tests.

### License
MIT © 2025