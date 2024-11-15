package backend

import (
	"BetterPC_2.0/pkg/logging"
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler, logger *logging.Logger) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20, //1MB
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	logger.Infof("Server started on port %s", port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
