package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"month02blogapi/config"
	"month02blogapi/handler"
	"month02blogapi/repository"
	"month02blogapi/service"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appConfig, err := config.Load("config.json")
	if err != nil {
		logger.Error("load config failed", "error", err)
		return
	}

	postRepository := repository.NewInMemoryPostRepository()
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService, logger.With("component", "post_handler"))

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	postHandler.RegisterRoutes(mux)

	logger.Info("blog api starting", "address", appConfig.Address())
	if err := http.ListenAndServe(appConfig.Address(), mux); err != nil {
		logger.Error("server stopped", "error", err)
	}
}
