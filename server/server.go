package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Server is a http server.
type Server struct {
	port   int
	logger *logrus.Logger
	server *http.Server
}

// NewServer is a constructor.
func NewServer(logger *logrus.Logger, port int, handler http.Handler) *Server {
	server := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 60,
	}

	return &Server{
		port:   port,
		logger: logger,
		server: server,
	}
}

// ListenAndServe will start the http server. Do not call it with goroutine.
func (s Server) ListenAndServe() {
	go func() {
		s.logger.Infof("Server is running on port :%d", s.port)
		s.logger.Error(s.server.ListenAndServe())
	}()
}

// Shutdown will gracefully shutdown the http server.
func (s Server) Shutdown() {
	s.logger.Error(s.server.Shutdown(context.Background()))
}
