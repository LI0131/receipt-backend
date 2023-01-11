package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/li0131/receipt-backend/internal/config"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func main() {
	// Get the configuration
	cfg := config.Get()

	// Create the routers
	rootRouter := chi.NewRouter()
	endpointsRouter := chi.NewRouter()
	rootRouter.Use(middleware.Logger)
	rootRouter.Mount(cfg.BaseApiPath, endpointsRouter)

	// Apply the endpoint handlers
	endpointsRouter.Get("/", hello)

	// create server and serve
	srv := http.Server{
		Addr:    cfg.Hostname + ":" + cfg.PublicPort,
		Handler: rootRouter,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
