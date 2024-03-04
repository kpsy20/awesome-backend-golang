package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/ztrue/shutdown"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	Host                string   `envDefault:"127.0.0.1"`
	Port                int      `envDefault:"8080"`
	LoggingPathPrefixes []string `envDefault:""`

	root   *chi.Mux
	server *http.Server

	running context.Context
	cancel  context.CancelFunc
}

func NewServer(config *Server) (s *Server) {
	root := chi.NewRouter()

	s = &Server{
		root:                root,
		Host:                config.Host,
		Port:                config.Port,
		LoggingPathPrefixes: config.LoggingPathPrefixes,
	}

	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler:      root,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return
}

func (s *Server) Mount(path string, handler http.Handler) {
	s.root.Mount(path, handler)
}

func (s *Server) Start() {
	s.running, s.cancel = context.WithCancel(context.Background())
	s.printRoutePath()

	shutdown.Add(func() { s.Stop() })

	// Run the server
	slog.Debug("listen on", "host", s.Host, "port", s.Port)

	go func() {
		if err := s.server.ListenAndServe(); nil != err && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("serve", "error", err)
			s.Stop()
			panic(err)
		}
	}()
}

func (s *Server) Stop() {
	stop, cancel := context.WithTimeout(s.running, 30*time.Second)
	go func() {
		<-stop.Done()
		slog.Debug("graceful stop")
		if errors.Is(stop.Err(), context.DeadlineExceeded) {
			slog.Warn("graceful stop, timeout", "error", stop.Err())
			cancel()
		}
	}()

	slog.Debug("shutdown")
	if err := s.server.Shutdown(stop); err != nil {
		slog.Warn("shutdown", "error", err)
	}
	s.cancel()
}

func (s *Server) printRoutePath() {
	routes := make([]string, 0)
	f := func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) (err error) {
		routes = append(routes, fmt.Sprintf("%s %s", method, route))
		return
	}
	_ = chi.Walk(s.root, f)
	slog.Debug("serve routes", "count", len(routes), "path", routes)
}
