# 📝 Blog Project

A modern, RESTful Blog API built with Go, Echo framework, and PostgreSQL. Features JWT authentication, CRUD operations for posts, comments, and reactions with a clean architecture following repository pattern.

---

## 🌟 Features

- **User Authentication** — Secure signup/login with JWT tokens and bcrypt password hashing
- **Blog Posts** — Create, read, update, and delete posts with slug-based URLs
- **Comments** — Add and list comments on posts
- **Reactions** — Toggle reactions (like/unlike) on posts
- **Tags** — Tag-based categorization for posts
- **Role-based Access** — Reader/Author role system
- **Clean Architecture** — Repository → Service → Handler pattern
- **UUID Primary Keys** — Secure, non-sequential identifiers
- **Auto Migration** — Database schema auto-synced on startup

---

## 🛠️ Tech Stack

| Layer | Technology |
|-------|-----------|
| **Language** | Go 1.25+ |
| **Web Framework** | [Echo v4](https://echo.labstack.com/) |
| **ORM** | [GORM](https://gorm.io/) |
| **Database** | PostgreSQL |
| **Authentication** | JWT (golang-jwt) |
| **Password Hashing** | bcrypt (golang.org/x/crypto) |
| **Env Management** | godotenv |
| **UUID** | google/uuid |

---

## 📁 Project Structure

```
blog-project/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── config/
│   └── config.go                # Environment configuration loader
├── internal/
│   ├── app/
│   │   ├── server.go            # Echo server initialization
│   │   ├── routes.go            # API route definitions
│   │   └── middleware.go        # JWT & DB injection middleware
│   ├── auth/
│   │   ├── jwt.go               # JWT token generation & validation
│   │   └── password.go          # Password hashing utilities
│   ├── database/
│   │   └── db.go                # PostgreSQL connection manager
│   ├── handler/
│   │   ├── auth_handlers.go     # Signup & login handlers
│   │   ├── post_handlers.go     # Post CRUD handlers
│   │   ├── comment_handlers.go  # Comment handlers
│   │   └── reaction_handlers.go # Reaction toggle handler
│   ├── models/
│   │   ├── user.go              # User model
│   │   ├── post.go              # Post model
│   │   ├── comment.go           # Comment model
│   │   ├── reaction.go          # Reaction model
│   │   └── tag.go               # Tag model
│   ├── repositories/
│   │   ├── user_repository.go   # User DB operations
│   │   ├── post_repository.go   # Post DB operations
│   │   ├── comment_repositiry.go # Comment DB operations
│   │   └── reaction_repository.go # Reaction DB operations
│   ├── services/
│   │   ├── user_service.go      # User business logic
│   │   ├── post_service.go      # Post business logic
│   │   ├── comment_service.go   # Comment business logic
│   │   └── reaction_service.go  # Reaction business logic
│   └── utils/
│       ├── response.go          # Standard JSON response helper
│       └── slug.go              # URL slug generator
├── .env.sample                  # Environment variables template
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
└── README.md                    # This file
```

---

## 🚀 Getting Started

### Prerequisites

- **Go** 1.25 or higher
- **PostgreSQL** running instance (local or cloud like [Neon](https://console.neon.tech/))

### 1. Clone the repository

```bash
git clone https://github.com/surajit/blog-project.git
cd blog-project
```

### 2. Set up environment variables

```bash
cp .env.sample .env
```

Edit `.env` with your database credentials:

```env
DATABASE_URL="postgres://username:password@localhost:5432/blog_project?sslmode=disable"
APP_PORT=8080
JWT_SECRET="your-super-secret-key"
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the server

```bash
go run cmd/server/main.go
```

The server starts at `http://localhost:8080`

---

## 📡 API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/signup` | Register a new user |
| `POST` | `/api/v1/login` | Login and receive JWT |
| `GET` | `/api/v1/posts` | List all posts |
| `GET` | `/api/v1/posts/:id` | Get post by ID |
| `GET` | `/api/v1/posts/:id/comments` | List comments for a post |
| `GET` | `/api/v1/tags` | List all tags |

### Protected Endpoints (Requires JWT Bearer Token)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/posts` | Create a new post |
| `PUT` | `/api/v1/posts/:id` | Update a post |
| `DELETE` | `/api/v1/posts/:id` | Delete a post |
| `POST` | `/api/v1/posts/:id/comments` | Add a comment to a post |
| `POST` | `/api/v1/posts/:id/reactions` | Toggle reaction on a post |

---

## 🔐 Authentication

Include the JWT token in the `Authorization` header for protected routes:

```
Authorization: Bearer <your-jwt-token>
```

---

## 📊 Data Models

### User
| Field | Type | Description |
|-------|------|-------------|
| `id` | UUID | Primary key |
| `username` | string | Unique username |
| `email` | string | Unique email |
| `password` | string | Hashed password (hidden in JSON) |
| `display_name` | string | Display name |
| `bio` | string | User biography |
| `role` | string | `reader` (default) or `author` |
| `is_active` | bool | Account status |

### Post
| Field | Type | Description |
|-------|------|-------------|
| `id` | UUID | Primary key |
| `author_id` | UUID | Foreign key to User |
| `title` | string | Post title |
| `slug` | string | URL-friendly slug |
| `excerpt` | string | Short summary |
| `content` | text | Full post content |
| `status` | string | `draft` (default) or `published` |
| `visibility` | string | `public` (default) or `private` |
| `views` | int64 | View counter |
| `tags` | []Tag | Associated tags |

---

## 🗄️ Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DATABASE_URL` | `postgres://postgres:postgres@localhost:5432/blog_project?sslmode=disable` | PostgreSQL connection string |
| `APP_PORT` | `8080` | Server listening port |
| `JWT_SECRET` | `secret` | Secret key for JWT signing |

---

## 📂 Architecture

```
Request → Handler → Service → Repository → Database
   ↑                                        |
   └────────────── Response ←───────────────┘
```

- **Handlers** — HTTP request/response handling, input validation
- **Services** — Business logic layer
- **Repositories** — Database operations (GORM queries)
- **Models** — Data structures and relationships

---

## 🤝 Contributing

Contributions are welcome! Please read our guidelines before getting started:

- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Contributing Guidelines](CONTRIBUTING.md)

### Quick Start

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

**Built with ❤️ using Go & Echo**
