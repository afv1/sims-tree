package sims_tree

import (
    "context"
    "github.com/afv1/sims-tree/cfg"
    "net/http"
    "time"
)

type Server struct {
    httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
    s.httpServer = &http.Server{
        Addr: cfg.App.Host + ":" + cfg.App.Port,
        Handler: handler,
        MaxHeaderBytes: 1 << 20,
        ReadTimeout: time.Duration(cfg.App.ReadTimeout) * time.Second,
        WriteTimeout: time.Duration(cfg.App.WriteTimeout) * time.Second,
    }

    return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
    return s.httpServer.Shutdown(ctx)
}
