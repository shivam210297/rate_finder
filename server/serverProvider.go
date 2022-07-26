package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultServerRequestTimeoutMinutes      = 2
	defaultServerReadHeaderTimeoutSeconds   = 30
	defaultServerRequestWriteTimeoutMinutes = 30
)

type Server struct {
	httpServer *http.Server
}

func SrvInit() *Server {
	return &Server{}
}

func (srv *Server) Start() {

	addr := ":8080"
	httpSrv := &http.Server{
		Addr:              addr,
		Handler:           srv.InjectRoutes(),
		ReadTimeout:       defaultServerRequestTimeoutMinutes * time.Minute,
		ReadHeaderTimeout: defaultServerReadHeaderTimeoutSeconds * time.Second,
		WriteTimeout:      defaultServerRequestWriteTimeoutMinutes * time.Minute,
	}
	srv.httpServer = httpSrv

	logrus.Info("Server running at PORT ", addr)
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err)
		return
	}
}

func (srv *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logrus.Info("closing server...")
	_ = srv.httpServer.Shutdown(ctx)
	logrus.Info("Done")
}
