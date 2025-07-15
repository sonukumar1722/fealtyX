
# 🎓 Student Management API with AI Summary (Go + Ollama)

A simple, blazing-fast RESTful API built with Go for managing student records. Supports single & bulk CRUD operations and integrates with **Ollama** to generate professional AI summaries of student profiles.

---

## 🚀 Features

✅ Add single or multiple students  
✅ View all students or a specific student  
✅ Update one or more students  
✅ Delete single or multiple students  
✅ Generate AI-based summaries using Ollama (`llama3`)  
✅ Deployed on Render with redirect from `/` → `/students`

---

## 🌐 Live API URL

🔗 [https://student-api-bez3.onrender.com/students](https://student-api-bez3.onrender.com/students)

> ℹ️ The base URL `/` automatically redirects to `/students`.

---

## ⚙️ Requirements

- [Go 1.18+](https://golang.org/doc/install)
- [Ollama](https://ollama.com/) installed locally **(for AI Summary feature)**

```bash
ollama serve
ollama pull llama3
```

---

## 📁 Project Structure

```
.
├── main.go           # Entry point and route definitions
├── student.go        # All CRUD logic and in-memory storage
├── ollama.go         # AI summary generation using Ollama API
├── render.yaml       # Render deployment config
├── go.mod            # Go module metadata
└── README.md         # This file
```

---

## 🧪 Run the Server Locally

```bash
go run .
```

> Server runs at: `http://localhost:8080`

---

## 📦 API Endpoints

| Method  | Endpoint                    | Description                       |
|---------|-----------------------------|-----------------------------------|
| POST    | `/students`                 | Add one or more students          |
| GET     | `/students`                 | Get all students                  |
| GET     | `/students/{id}`           | Get a student by ID               |
| PUT     | `/students`                 | Bulk update students              |
| PUT     | `/students/{id}`           | Update a specific student         |
| DELETE  | `/students?ids=1,2`         | Bulk delete students              |
| DELETE  | `/students/{id}`           | Delete one student by ID          |
| GET     | `/students/{id}/summary`   | Generate AI summary (local only)  |

> Note: `/summary` endpoint returns a fallback message on deployed URL since Ollama runs only locally.

---

## 🧪 Sample `curl` Tests

### ➕ Create Students
```bash
curl -X POST https://student-api-bez3.onrender.com/students -H "Content-Type: application/json" -d '[{"id":1,"name":"Alice","age":20,"email":"alice@example.com"}]'
```

### 📖 Get All Students
```bash
curl https://student-api-bez3.onrender.com/students
```

### ✏️ Update a Student
```bash
curl -X PUT https://student-api-bez3.onrender.com/students/1 -H "Content-Type: application/json" -d '{"name":"Alice Smith","age":21,"email":"alice.smith@example.com"}'
```

### ❌ Delete a Student
```bash
curl -X DELETE https://student-api-bez3.onrender.com/students/1
```

### 🧹 Bulk Delete
```bash
curl -X DELETE "https://student-api-bez3.onrender.com/students?ids=1,2"
```

### 🤖 Get AI Summary (Local only)
```bash
curl http://localhost:8080/students/1/summary
```

---

## 🤖 Ollama Setup (Local Only)

Make sure Ollama is running locally:

```bash
ollama serve
ollama pull llama3
```

The app uses the `/api/chat` endpoint for generating summaries.

---

## 💡 TODO / Suggestions

- 🔒 Add authentication
- 🧠 Support more LLMs (e.g. Mistral, Codellama)
- 🐳 Add Docker support
- 📄 Add OpenAPI / Swagger spec
- 🛢️ Replace in-memory storage with PostgreSQL

---

## 👨‍💻 Author

Built with ❤️ by **Sonu Kumar Shah**  
[GitHub](https://github.com/sonukumar1722) • [LinkedIn](https://linkedin.com/in/sonukumar287)

---

## 📜 License

MIT License. Free to use and modify.
