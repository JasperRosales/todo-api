# Todo RestFul API

A practice api for testing my backend skills only using the documentation while reading the project guidelines from roadmap.sh. 


project url: https://roadmap.sh/projects/todo-list-api
---

## Requirements: 
- User registration to create a new user
- Login endpoint to authenticate the user and generate a token
- CRUD operations for managing the to-do list
- Implement user authentication to allow only authorized users to access the to-do list
- Implement error handling and security measures
- Use a database to store the user and to-do list data (you can use any database of your choice)
- Implement proper data validation
- Implement pagination and filtering for the to-do list

---

### Schema

The data schema for each models to be used.


User Model:
```sql
CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL, 
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL
)
```

Todo Model:
```sql
CREATE TABLE IF NOT EXISTS todos(
    todo_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    user_id INT,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE
)
```

