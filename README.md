# Go REST API Example

A simple RESTful API for event management built with Go, Gin, and SQLite. This project demonstrates user authentication (JWT), event CRUD operations, and event registration/cancellation.

## Features

- User signup and login with password hashing (bcrypt)
- JWT-based authentication middleware
- CRUD operations for events
- Event registration and cancellation
- SQLite database with auto-migration
- Organized code structure (models, routes, middlewares, utils)

## Tech Stack

- [Go](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [SQLite](https://www.sqlite.org/)
- [JWT](https://github.com/golang-jwt/jwt)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## Getting Started

### Prerequisites

- Go 1.18+
- SQLite3

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/go-rest-api.git
    cd go-rest-api
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

The API will be available at `http://localhost:8080`.

## API Endpoints

### Public

- `POST /signup` — Register a new user
- `POST /login` — Login and receive JWT
- `GET /events` — List all events
- `GET /events/:id` — Get event by ID

### Protected (require JWT in `Authorization` header)

- `POST /events` — Create event
- `PUT /events/:id` — Update event (owner only)
- `DELETE /events/:id` — Delete event (owner only)
- `POST /events/:id/register` — Register for event
- `DELETE /events/:id/register` — Cancel registration

## Example Requests

See the `api-test/` directory for sample HTTP requests using [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

## Project Structure

```
.
├── api-test/         # Sample HTTP requests
├── db/               # Database initialization
├── middlewares/      # JWT authentication middleware
├── models/           # Database models
├── routes/           # API route handlers
├── utils/            # Utility functions (hashing, JWT)
├── main.go           # Entry point
└── go.mod, go.sum    # Go modules
```

## Database

- Uses SQLite (`api.db`).
- Tables: `users`, `events`, `registrations`.
- Tables are auto-created on startup.

## Security

- Passwords are hashed with bcrypt.
- JWT tokens are used for authentication and must be sent in the `Authorization` header.

---

**Note:** This project is for educational/demo purposes. For production, use environment variables for secrets, add input validation, and follow security best practices.