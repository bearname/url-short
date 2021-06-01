package infrastructure

import (
	"github.com/bearname/url-short/pkg/short/infrastructure/router"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
}

func (s *Server) StartServer(serverUrl string, client *mongo.Client, collection *mongo.Collection) *http.Server {
	router := router.Router(client, collection)
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Error(srv.ListenAndServe())
	}()

	return srv
}

func (s *Server) GetKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func (s *Server) WaitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
