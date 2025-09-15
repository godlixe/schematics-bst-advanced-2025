package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Struct Blog memodelkan blog pada program.
type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}

// Struct Response memodelkan response API.
type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

var Database []Blog

// POST /blogs
// Digunakan untuk membuat entry blog.
// Handler ini menerima atribut title, content, dan author.
// Atribut ID dan timestamp akan ditambahkan oleh program.
func createBlogsHandler(w http.ResponseWriter, r *http.Request) {
	var blog Blog

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// tambah id dan timestamp
	blog.ID = len(Database) + 1
	blog.Timestamp = time.Now()

	Database = append(Database, blog)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(blog)
}

// GET /blogs
// Digunakan untuk mendapatkan semua entry blog.
func getAllBlogsHandler(w http.ResponseWriter, r *http.Request) {
	var response = Response{
		Data:    Database,
		Message: "get all blogs successful",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET /blogs/{id}
// Digunakan untuk mendapatkan salah satu entry blog dengan ID tertentu.
func getBlogByIDHandler(w http.ResponseWriter, r *http.Request) {
	// mendapatkan ID dari path
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid blog id", http.StatusBadRequest)
		return
	}

	// cari blog dengan ID tertentu
	for _, blog := range Database {
		if blog.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(blog)
			return
		}
	}

	http.Error(w, "blog not found", http.StatusNotFound)
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /blogs", createBlogsHandler)
	mux.HandleFunc("GET /blogs", getAllBlogsHandler)
	mux.HandleFunc("GET /blogs/{id}", getBlogByIDHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)

}
