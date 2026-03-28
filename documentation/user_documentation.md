# User Documentation

## Introduction
This document provides guidance on managing users using the API. Supports user registration, login, CRUD operations. Passwords are hashed. JWT tokens issued on login (24h expiry).

## API Endpoints
- POST /users: Create a new user.
- POST /login: Login and get JWT token.
- GET /users/:id: Retrieve a specific user.
- PUT /users/:id: Update a user (partial updates supported, including password).
- DELETE /users/:id: Delete a specific user.
- GET /users: List users with pagination (page and limit query params).

Update/Delete own user requires auth; admin for others (assumed).

## Examples

### Create User
Request:
```
POST /users
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepass123"
}
```

### Login
Request:
```
POST /login
{
  "email": "john@example.com",
  "password": "securepass123"
}
```
Response includes token.

### List Users
Request:
```
GET /users?page=1&amp;limit=10
```

Response:
```
{
  "data": [...],
  "total": 25
}
```

### Update User
Request:
```
PUT /users/1
{
  "name": "John Doe Updated",
  "password": "newpass456"
}
```

## Common Tasks
- Register: Provide name, email (unique), password (min 8 chars).
- Login: Get token for auth.
- Update partially: Send name, email, or password.
- Pagination: Default page=1, limit=10 (max 100).
- Password updates: Hashed automatically.

Note: Email must be valid. Password min length 8.

## Rate Limiting
All endpoints are rate limited to 100 requests per minute per IP address. Exceeding the limit returns HTTP 429 Too Many Requests. Headers include rate limit info.
