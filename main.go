package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"service1/server"
	"syscall"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logrus.Info(os.Getenv("PORT"))
	srv := server.SrvInit()
	go srv.Start()

	<-done
	logrus.Info("Graceful shutdown")
	srv.Stop()
}
