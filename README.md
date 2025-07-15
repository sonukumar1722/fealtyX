
# 🎓 Student Management API with AI Summary (Go + Ollama)

A simple, blazing-fast RESTful API built with Go for managing student records. Supports single & bulk CRUD operations and integrates with **Ollama** to generate professional AI summaries of student profiles.

---

## 🚀 Features

✅ Add single or multiple students  
✅ View all students or a specific student  
✅ Update one or more students  
✅ Delete single or multiple students  
✅ Generate AI-based summaries using Ollama (`llama3`)

---

## 📁 Project Structure

```
.
├── main.go           # Entry point and route definitions
├── student.go        # All CRUD logic and in-memory storage
├── ollama.go         # AI summary generation using Ollama API
└── go.mod            # Go module metadata
```

---

## ⚙️ Requirements

- [Go 1.18+](https://golang.org/doc/install)
- [Ollama](https://ollama.com/) installed locally and running:
```bash
  ollama serve
  ollama pull llama3
```

---

## 🧪 Run the Server

```bash
go run .
```

Server will start on:
```
http://localhost:8080
```

---

## 📦 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/students` | Add one or more students |
| `GET` | `/students` | Get all students |
| `GET` | `/students/{id}` | Get a student by ID |
| `PUT` | `/students` | Bulk update students |
| `PUT` | `/students/{id}` | Update a specific student |
| `DELETE` | `/students?ids=1,2` | Bulk delete |
| `DELETE` | `/students/{id}` | Delete one student |
| `GET` | `/students/{id}/summary` | Generate AI summary via Ollama |

---

## 🧪 Sample `curl` Tests

### ➕ Create Students
```bash
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '[
      {"id":1,"name":"Alice","age":20,"email":"alice@example.com"}
    ]'
```

### 📖 Get All Students
```bash
curl http://localhost:8080/students
```

### ✏️ Update a Student
```bash
curl -X PUT http://localhost:8080/students/1 \
-H "Content-Type: application/json" \
-d '{"name":"Alice Smith","age":21,"email":"alice.smith@example.com"}'
```
OR
```bash
curl -X PUT http://localhost:8080/students/1 \
-H "Content-Type: application/json" \
-d '[
      {"name":"Alice Smith","age":21,"email":"alice.smith@example.com"}, {"name":"Alice Smith","age":21,"email":"alice.smith@example.com"}
    ]'
```

### ❌ Delete a Student
```bash
curl -X DELETE http://localhost:8080/students/1
```
OR
```bash
curl -X DELETE "http://localhost:8080/students?ids=1,2"
```

### 🤖 Get AI Summary
```bash
curl http://localhost:8080/students/1/summary
```

---

## 🤖 Ollama Setup

Make sure Ollama is running locally:

```bash
ollama serve
```

Pull a model (e.g., `llama3`):

```bash
ollama pull llama3
```

> The AI summary uses the `/api/chat` endpoint with the `llama3` model.

---

<!-- ## 📌 TODO / Suggestions

- 🔒 Add authentication
- 🧠 Support more LLMs via Ollama (Mistral, Codellama, etc.)
- 🐳 Dockerize the app
- 📄 Add Swagger / OpenAPI spec
- 🛢️ Replace in-memory DB with PostgreSQL or SQLite

--- -->

## 🧑‍💻 Author

Built with ❤️ by Sonu Kumar Shah  
[GitHub](https://github.com/sonukumar1722) | [LinkedIn](https://linkedin.com/in/sonukumar287)

---

## 📜 License

MIT License. Free to use and modify.
# fealtyX
