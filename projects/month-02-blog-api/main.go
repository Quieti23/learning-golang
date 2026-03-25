package main

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"month02blogapi/config"
	"month02blogapi/handler"
	"month02blogapi/middleware"
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

	db, err := sql.Open("mysql", appConfig.MySQLDSN)
	if err != nil {
		logger.Error("open mysql failed", "error", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error("ping mysql failed", "error", err)
		return
	}

	db.SetMaxOpenConns(appConfig.MaxOpenConns)
	db.SetMaxIdleConns(appConfig.MaxIdleConns)
	db.SetConnMaxLifetime(appConfig.ConnMaxLifetime())

	postRepository := repository.NewMySQLPostRepository(db)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService, logger.With("component", "post_handler"), appConfig.RequestTimeout())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	postHandler.RegisterRoutes(mux)
	serverHandler := middleware.RequestLogger(logger.With("component", "http_middleware"))(mux)

	logger.Info(
		"blog api starting",
		"address", appConfig.Address(),
		"database", "mysql",
		"max_open_conns", appConfig.MaxOpenConns,
		"max_idle_conns", appConfig.MaxIdleConns,
		"conn_max_lifetime_minutes", appConfig.ConnMaxLifetimeMins,
		"request_timeout_ms", appConfig.RequestTimeoutMs,
	)
	if err := http.ListenAndServe(appConfig.Address(), serverHandler); err != nil {
		logger.Error("server stopped", "error", err)
	}
}
