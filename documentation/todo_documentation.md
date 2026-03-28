# Todo Documentation

## Introduction
This document provides guidance on managing todos using the API. Todos are user-scoped, meaning each todo belongs to a specific user via authentication.

## API Endpoints
- POST /todos: Create a new todo.
- GET /todos/:id: Retrieve a specific todo.
- PUT /todos/:id: Update a todo (partial updates supported).
- DELETE /todos/:id: Delete a specific todo.
- GET /todos: List todos with pagination (page and limit query params).

All endpoints require authentication. Todos are filtered by the authenticated user ID.

## Examples

### Create Todo
Request:
```
POST /todos
{
  "title": "Buy groceries",
  "description": "Milk, bread, eggs"
}
```

### List Todos
Request:
```
GET /todos?page=1&amp;limit=10
```

Response:
```
{
  "data": [...],
  "total": 25
}
```

### Update Todo
Request:
```
PUT /todos/1
{
  "title": "Buy groceries (updated)"
}
```

## Common Tasks
- List all your todos: Use GET /todos.
- Create a todo: Provide title and description.
- Update partially: Send only changed fields (title or description).
- Delete: Specify the todo ID.
- Pagination: Default page=1, limit=10 (max 100).

Note: Ensure valid authentication token in requests. Title and description are required for creation (min length 1).

## Rate Limiting
All endpoints are rate limited to 100 requests per minute per IP address. Exceeding the limit returns HTTP 429 Too Many Requests. Headers include rate limit info.
