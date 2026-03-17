package webhook

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/webhook/api"
)

type Options struct {
	Host         string        `env:"HOST, default=localhost"`
	Port         string        `env:"PORT, default=8888"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT, default=60s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT, default=60s"`
}

func (o *Options) Addr() string {
	return fmt.Sprintf("%s:%s", o.Host, o.Port)
}

type Server struct {
	Provider    provider.Provider
	Options     Options
	HealthyFunc func(bool)
}

func (s *Server) Serve(ctx context.Context) error {
	if s.Provider == nil {
		return fmt.Errorf("provider is required")
	}

	p := api.WebhookServer{
		Provider: s.Provider,
	}

	m := http.NewServeMux()
	m.HandleFunc("/", p.NegotiateHandler)
	m.HandleFunc("/records", p.RecordsHandler)
	m.HandleFunc("/adjustendpoints", p.AdjustEndpointsHandler)

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

		s.setHealthy(false)

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}

		if err := l.Close(); err != nil { // Check the error here
			log.Fatal(err)
		}
	}()

	s.setHealthy(true)

	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Server) setHealthy(healthy bool) {
	if s.HealthyFunc == nil {
		return
	}

	s.HealthyFunc(healthy)
}
