package main

import (
	"awesome-backend-golang/internal/adapter/api"
	"awesome-backend-golang/internal/adapter/repository"
	"awesome-backend-golang/internal/adapter/repository/sqlite"
	"awesome-backend-golang/internal/aggregate/rooms"
	"awesome-backend-golang/pkg/env"
	"awesome-backend-golang/pkg/http"
	"github.com/ztrue/shutdown"
	"log/slog"
	"syscall"
)

func main() {
	env.MustLoadEnvDefault()
	slog.Info("🔫 server start")
	db := sqlite.NewSqliteDB(env.MustFetchEnv(&sqlite.Config{}, "SQLITE_"))

	repo := repository.New(db)

	room := rooms.New(repo)

	rest := api.New(api.WithRooms(room))

	server := http.NewServer(env.MustFetchEnv(&http.Server{}, "HTTP_"))
	server.Mount("/v1", rest.NewRouter())
	server.Start()
	shutdown.Listen(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
}
