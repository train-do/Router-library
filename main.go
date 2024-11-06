package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/train-do/Router-library/database"
	"github.com/train-do/Router-library/handler"
	mid "github.com/train-do/Router-library/middleware"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Group(func(r chi.Router) {
		r.Get("/login", handler.Login(db))
		r.Post("/login", handler.Login(db))
		r.Get("/users", handler.GetUsers(db))
		r.Get("/register", handler.Register(db))
		r.Post("/register", handler.Register(db))
	})

	router.Group(func(r chi.Router) {
		r.Use(mid.Authentication)
		r.Get("/todo/all", handler.GetTodo(db))
		r.Post("/todo/create", handler.CreateTodo(db))
	})

	fmt.Println("server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
