package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/li0131/receipt-backend/internal/api"
	"github.com/li0131/receipt-backend/internal/config"
)

func hello(w http.ResponseWriter, r *http.Request) {
	api.Success(w, "Hello World")
}

func main() {
	// Get the configuration
	cfg := config.Get()

	// Create the routers
	rootRouter := chi.NewRouter()
	endpointsRouter := chi.NewRouter()
	rootRouter.Use(middleware.Logger)
	rootRouter.Mount(cfg.BaseApiPath, endpointsRouter)

	// Add error handlers
	rootRouter.NotFound(api.NotFound)
	rootRouter.MethodNotAllowed(api.MethodNotAllowed)

	// Apply the endpoint handlers
	endpointsRouter.Get("/", hello)
	endpointsRouter.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		api.UploadFile(cfg.FileServerConfig, w, r)
	})

	// create server and serve
	srv := http.Server{
		Addr:    cfg.Hostname + ":" + cfg.PublicPort,
		Handler: rootRouter,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
