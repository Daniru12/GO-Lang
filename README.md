
# Simple CRUD API in Go

## Overview

This is a **simple CRUD (Create, Read, Update, Delete) application** built with **Golang** and **Gorilla Mux**.
The project is designed to help understand **RESTful APIs** and practice backend development.

* **Language:** Go
* **Framework:** Gorilla Mux
* **Purpose:** Learn API design, routing, and basic CRUD operations

---

## Features

* **Create** new tasks
* **Read** all tasks or a single task by ID
* **Update** existing tasks
* **Delete/Update status** of tasks
* Clean and simple structure for learning purposes

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/simple-crud-api.git
```

2. Navigate to the project directory:

```bash
cd simple-crud-api
```

3. Run the application:

```bash
go run main.go
```

The server will start at `http://localhost:8080` by default.

---

## API Endpoints

| Method | Endpoint                      | Description                        |
| ------ | ----------------------------- | ---------------------------------- |
| POST   | `/tasks`                      | Create a new task                  |
| GET    | `/tasks`                      | Get all tasks                      |
| GET    | `/tasks/{resource_id}`        | Get a task by its resource ID      |
| PATCH  | `/tasks/{resource_id}`        | Update a task’s details            |
| PATCH  | `/tasks/{resource_id}/status` | Update task status or mark deleted |

**Notes:**

* `{resource_id}` is a **path parameter** used to identify a specific task.
* The API follows **REST principles** for simplicity and scalability.

---

## Usage

You can test the API using **Postman**, **curl**, or any API testing tool.

**Example using curl:**

* Create a task:

```bash
curl -X POST http://localhost:8080/tasks -d '{"title":"My Task","description":"This is a test"}' -H "Content-Type: application/json"
```

* Get all tasks:

```bash
curl http://localhost:8080/tasks
```

* Get a task by ID:

```bash
curl http://localhost:8080/tasks/1
```

* Update a task:

```bash
curl -X PATCH http://localhost:8080/tasks/1 -d '{"title":"Updated Task"}' -H "Content-Type: application/json"
```

* Update task status:

```bash
curl -X PATCH http://localhost:8080/tasks/1/status -d '{"status":"completed"}' -H "Content-Type: application/json"
```

---

## Project Structure

```
simple-crud-api/
│
├─ main.go           # Entry point of the application
├─ transport/
│   └─ router.go     # API routes using Gorilla Mux
├─ endpoints/
│   └─ task_handler.go  # Handlers for API endpoints
└─ models/           # Optional: data models for tasks
```

---

## Contributing

* Fork the repository, make your changes, and create a pull request.
* Feel free to experiment with features or improve code structure.

---

## License

This project is open source and available under the [MIT License](LICENSE).
