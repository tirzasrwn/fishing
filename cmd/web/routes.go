package main

import (
	"net/http"

	"github.com/tirzasrwn/fishing/internal/config"
	"github.com/tirzasrwn/fishing/internal/handler"
)

func routes(app *config.AppConfig) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.Repo.Home)
	mux.HandleFunc("POST /", handler.Repo.PostHome)
	mux.HandleFunc("GET /success", handler.Repo.Success)
	mux.HandleFunc("GET /list", handler.Repo.List)
	return mux
}
