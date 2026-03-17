package health

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

type Options struct {
	Host         string        `env:"HOST, default=0.0.0.0"`
	Port         string        `env:"PORT, default=8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT, default=60s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT, default=60s"`
}

func (o *Options) Addr() string {
	return fmt.Sprintf("%s:%s", o.Host, o.Port)
}

type Server struct {
	Options Options
	healthy atomic.Bool
}

func (s *Server) SetHealthy(healthy bool) {
	switch healthy {
	case true:
		slog.Info("Service is now healthy.")
	default:
		slog.Warn("Service is unhealthy.")
	}

	s.healthy.Store(healthy)
}

func (s *Server) Healthy() bool {
	return s.healthy.Load()
}

func (s *Server) Serve(ctx context.Context) error {
	m := http.NewServeMux()
	m.HandleFunc("/healthz", s.handleHealthz)

	srv := &http.Server{
		Addr:         s.Options.Addr(),
		Handler:      m,
		ReadTimeout:  s.Options.ReadTimeout,
		WriteTimeout: s.Options.WriteTimeout,
	}

	l, err := net.Listen("tcp", s.Options.Addr())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		if err := l.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	if s.Healthy() {
		writeResponse(w, http.StatusOK, "Healthy")
		return
	}

	writeResponse(w, http.StatusServiceUnavailable, "Not Healthy")
}

func writeResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	if _, err := w.Write([]byte(message)); err != nil {
		slog.Error("Failed to write response", slog.Any("error", err))
	}
}
