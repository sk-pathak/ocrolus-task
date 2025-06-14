# Ocrolus Task – Article Management Backend Service

This project is a backend service built with Go, using Gin for the HTTP framework, sqlc for type-safe SQL query generation, Goose for database migrations, and PostgreSQL as the primary datastore. It allows users to perform CRUD operations on articles and tracks recently viewed articles per user.

---

## Features
- CRUD operations for articles
- User authentication using JWT
- Track & retrieve recently viewed articles per user
- Rate limiter
- Fully dockerized
- Database Migrations

---

## Getting Started

### Prerequisites
- Go 1.20+
- Postgresql
- Make
- sqlc
- goose
- docker & docker compose

**Make sure both `.env` file and `.secrets` file is present in the root directory. An example for each of these files is included**

### Local Development
Please ensure the prerequisites are installed in the system

1. Clone the repository
```bash
git clone https://github.com/sk-pathak/ocrolus-task.git
cd ocrolus-task
```
2. Create .env and .secrets file with all environment variables from the example. If using different database password, make sure to change in `DB_PASSWORD`, `GOOSE_DBSTRING` (just before the @ symbol) and in `.secrets` file

3. Run The Server
```bash
make run
```
Or build & execute
```bash
make exec
```

The server will start on http://localhost:8080

### Using Docker
```bash
docker compose up --build
```

---

### API Endpoints

| Method | Endpoint                                             | Description                  |
| ------ | ---------------------------------------------------- | ---------------------------- |
| POST   | `/register`                                          | Register user                |
| POST   | `/login`                                             | Authenticate user            |
| GET    | `/users`                                             | List all users               |
| GET    | `/users/:id`                                         | Get user by id               |
| GET    | `/users/me/recent-views`                             | List recently viewed         |
| GET    | `/users/me/articles?limit={limit}&offset={offset}`   | List my articles             |
| GET    | `/articles?limit={limit}&offset={offset}`            | List all articles            |
| POST   | `/articles`                                          | Create an article            |
| GET    | `/articles/:id`                                      | View single article          |
| PUT    | `/articles/:id`                                      | Update an article            |
| DELETE | `/articles/:id`                                      | Delete an article            |


All endpoints except for `/articles`, `/register` and `/login` are protected by JWT authentication. Make sure to register & log in user before making other calls to API. Both `/articles` and `/users/me/articles` are paginated

## Next Steps

- Refresh Tokens for JWT (*more secure*)
- Role Based Access Control (*finer control*)
- Unit Tests