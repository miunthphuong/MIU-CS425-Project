# Backend

Gin API for the MIU CS425 project.

## Features

- Gin HTTP server
- JWT authentication
- Separate auth and API routes for `supplier`, `admin`, and `customer`
- PostgreSQL 18 connection through GORM
- Redis connection for caching/session-ready workflows
- MinIO object storage with bucket initialization
- Request logging with Zap
- CORS configured for the frontend

## Run Locally

1. Copy the environment file:

   ```sh
   cp .env.example .env
   ```

2. Start PostgreSQL, Redis, and MinIO:

   ```sh
   docker compose up -d
   ```

3. Run the API:

   ```sh
   go run ./cmd/api
   ```

The API runs on `http://localhost:8080` by default.

MinIO console runs on `http://localhost:9001` with `minioadmin` / `minioadmin`.

## Main Endpoints

- `GET /api/v1/health`
- `POST /api/v1/admin/auth/register`
- `POST /api/v1/admin/auth/login`
- `GET /api/v1/admin/me`
- `GET /api/v1/admin/dashboard`
- `GET /api/v1/admin/users`
- `POST /api/v1/admin/users`
- `GET /api/v1/admin/users/:id`
- `PUT /api/v1/admin/users/:id`
- `DELETE /api/v1/admin/users/:id`
- `POST /api/v1/supplier/auth/register`
- `POST /api/v1/supplier/auth/login`
- `GET /api/v1/supplier/me`
- `GET /api/v1/supplier/dashboard`
- `POST /api/v1/customer/auth/register`
- `POST /api/v1/customer/auth/login`
- `GET /api/v1/customer/me`
- `GET /api/v1/customer/dashboard`

## Register Request

```json
{
  "name": "Customer User",
  "email": "customer@example.com",
  "password": "password123"
}
```

Use the role-specific register endpoint. For example, `POST /api/v1/customer/auth/register` creates a customer.

## Authenticated Requests

Send the JWT returned from register or login:

```http
Authorization: Bearer <token>
```
