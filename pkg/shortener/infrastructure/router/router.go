package router

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	//subrouter := router.PathPrefix("/api/v1").Subrouter()

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
