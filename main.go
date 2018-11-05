package main

import (
	"context"
	"log"
	"net/http"

	"gocity/handle/middlewares"

	"gocity/handle"

	"github.com/go-chi/chi"
	"gocity/lib"
)

func main() {
	router := chi.NewRouter()
	cache := lib.NewCache()
	storage, err := lib.NewGCS(context.Background())
	if err != nil {
		storage, err = lib.NewMemoryStorage()
		if err != nil {
			log.Fatal(err)
		}
	}

	corsMiddleware := middlewares.GetCors("*")
	router.Use(corsMiddleware.Handler)

	analyzer := handle.AnalyzerHandle{
		Cache:   cache,
		Storage: storage,
	}

	router.Get("/api", analyzer.Handler)
	router.Get("/health", handle.HealthCheck)

	log.Println("Server started at http://localhost:4000")
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Print(err)
	}
}
