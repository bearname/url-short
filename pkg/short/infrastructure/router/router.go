package router

import (
	"github.com/bearname/url-short/pkg/short/app"
	"github.com/bearname/url-short/pkg/short/infrastructure/middleware"
	"github.com/bearname/url-short/pkg/short/infrastructure/mongodb"
	"github.com/bearname/url-short/pkg/short/infrastructure/transport"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Router(client *mongo.Client, collection *mongo.Collection) http.Handler {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1/url").Subrouter()
	repository := postgres.New()
	service := app.NewUrlService(repository)
	controller := transport.NewUrlController(service)
	api.HandleFunc("", middleware.DecodeCreateUrlRequest(controller.Create())).Methods(http.MethodPost)
	api.HandleFunc("/{shortUrl}", controller.Redirect()).Methods(http.MethodGet)
	return logMiddleware(router)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}
