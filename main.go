package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Person is data model
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Initialize chi new router
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Handle root path
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		msg := make(map[string]string)
		msg["message"] = "Hello, World!"
		resp, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		w.Write(resp)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		person := &Person{}
		err := json.NewDecoder(r.Body).Decode(person)
		if err != nil {
			panic(err)
		}
		resp, err := json.Marshal(person)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	})

	// Handle /panic path
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// Handle /ping path
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// Create the server
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Cannot start server: %s\n", err)
	}
}
