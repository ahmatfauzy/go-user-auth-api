# Auth-API  
A clean-architecture **User Authentication API** built with **Go (Golang)**, **Gin**, **MySQL**, **Redis**, and **JWT**.  
Register, login, refresh token, logout — all secured with JWT & black-list on Redis.

---

## 🚀 Features

- ✅ **RESTful API Endpoints**  
  Base route: `/api/v1`  
  Includes: `/register`, `/login`, `/refresh`, `/logout`, `/profile`

- 🔐 **JWT-Based Stateless Authentication**  
  Secure access using JSON Web Tokens for protected routes.

- 🧹 **Token Blacklisting via Redis**  
  Implements logout by blacklisting access tokens in Redis.

- 🏗️ **Automatic Database Migration**  
  Schema auto-migration using GORM on app startup.

- 🧠 **Clean Architecture**  
  Domain-driven and layered code structure (handler → usecase → repository).

- 🛢️ **MySQL & Redis Ready**  
  Tested and integrated with MySQL and Redis out of the box.

---

## 📦 Tech Stack
| Layer          | Tech                         |
|----------------|------------------------------|
| HTTP Framework | Gin                          |
| ORM            | GORM                         |
| DB             | MySQL 8+                     |
| Cache          | Redis                        |
| Auth           | JWT                          |
| Environment    | `.env`                       |
| License        | MIT                          |

---

## 🛠️ Quick Start

### 1. Clone & Install
```bash
git clone https://github.com/ahmatfauzy/go-user-auth-api.git
cd auth-api
go mod tidy
```

### 2. Run App
```bash
go run cmd/main.go
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
└── README.md
```

### Contributing
We love PRs!
1. Fork → create feature branch → push → open PR.
2. Follow conventional commits & add tests.

### License
MIT © [ahmatfauzy](https://github.com/ahmatfauzy/go-user-auth-api?tab=MIT-1-ov-file) 2025
