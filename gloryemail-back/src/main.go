package main

import (
	"gloryemailback/src/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/test", db.TestDb)
	router.Route("/emails", func(r chi.Router) {
		r.Get("/", db.GetEmails)
		r.Get("/{query}", db.GetEmails)
	})

	http.ListenAndServe(":3000", router)
}
