package server

import (
	"context"
	"github.com/carinfinin/risk-assessor/internal/config"
	"net/http"
)

type Server struct {
	http.Server
}

func New(cfg *config.Config, router *Router) *Server {
	s := new(Server)
	s.Handler = router.Handler
	s.Addr = cfg.Addr
	return s
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
