
# Simple CRUD API in Go

## Overview

This is a **simple CRUD (Create, Read, Update, Delete) application** built with **Golang**, **MySQL**, and following **Clean Architecture principles**.
The project is designed to help understand **RESTful APIs**, **layered architecture**, and practice backend development.

**Tech Stack:**

* **Language:** Go 1.21+
* **Framework:** Gorilla Mux (for routing)
* **Database:** MySQL
* **Purpose:** Learn API design, routing, CRUD operations, and environment-based configuration

---

## Features

* **Create** new tasks
* **Read** all tasks or a single task by resource ID
* **Update** existing tasks
* **Soft Delete** / update status of tasks
* Clean and maintainable structure (`domain`, `usecases`, `repositories`, `services`, `transport`)
* Environment-based configuration using `.env` file

---

## Installation & Setup

### 1️⃣ Clone the repository

```bash
git clone https://github.com/your-username/task-management-api.git
cd task-management-api
```

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Setup MySQL database

* Install & run MySQL server.
* Create a database:

```sql
CREATE DATABASE taskDB;
```

* Create the `tasks` table:

```sql
CREATE TABLE tasks (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    resource_id VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status CHAR(1) DEFAULT 'A'
);
```

### 4️⃣ Configure environment variables

* Copy the example `.env` file:

```bash
cp .env.example .env
```

* Update `.env` with your database credentials:

```env
DB_USER=root
DB_PASS=yourpassword
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=taskDB
DB_LOC=Asia%2FColombo
APP_PORT=8081
```

> **Note:** `DB_LOC` should be URL-encoded (`Asia%2FColombo` → "Asia/Colombo").
> `.env` file should **not** be committed to GitHub. Keep `.env.example` for reference.

---

## Running the Application

```bash
go run main.go
```

If successful, you should see:

```
✅ Database connected successfully!
🚀 Server started at :8081
```

Visit the API in your browser or via Postman at [http://localhost:8081](http://localhost:8081)

---

## API Endpoints

| Method | Endpoint                      | Description                      |
| ------ | ----------------------------- | -------------------------------- |
| POST   | `/tasks`                      | Create a new task                |
| GET    | `/tasks`                      | Get all tasks                    |
| GET    | `/tasks/{resource_id}`        | Get task by resource ID          |
| PUT    | `/tasks/{resource_id}`        | Update task                      |
| PATCH  | `/tasks/{resource_id}/status` | Update task status (soft delete) |

---

## Testing

Run unit tests:

```bash
go test ./...
```

---

## License

This project is available under the [MIT License](LICENSE).

-
