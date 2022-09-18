package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/qna-page/qna-page/db"
	"github.com/qna-page/qna-page/resource/user"
)

func main() {
	db := db.ConnectDB()
	r := chi.NewRouter()

	// allowedCharsets := []string{"UTF-8"}

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Heartbeat("/health"))

	user := user.InitHandler(user.InitRepo(db))

	// Routes
	r.Get("/user", user.GetUsers)
	r.Get("/user/{id}", user.GetUser)

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
	}
}
