package router

import (
	"fmt"
	"github.com/bearname/url-short/internal/short/infrastructure/transport"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	runtimepprof "runtime/pprof"
)

func Router(controller *transport.UrlController) http.Handler {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	api := router.PathPrefix("/api/v1/urls").Subrouter()

	router.HandleFunc("/health", healthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/ready", readyCheckHandler).Methods(http.MethodGet)
	api.HandleFunc("", controller.Create()).Methods(http.MethodPost)
	//api.HandleFunc("", middleware.DecodeCreateUrlRequest(controller.Create())).Methods(http.MethodPost)
	api.HandleFunc("/{shortUrl}", controller.Redirect()).Methods(http.MethodGet)

	pprofRouter := router.PathPrefix("/debug/pprof").Subrouter()
	pprofRouter.HandleFunc("/", pprof.Index)
	pprofRouter.HandleFunc("/cmdline", pprof.Cmdline)
	pprofRouter.HandleFunc("/symbol", pprof.Symbol)
	pprofRouter.HandleFunc("/trace", pprof.Trace)

	// Debug: pprof.WriteHeapProfile()
	memprofile := "/run/dnsd/memprofile"
	f, _ := os.Create(memprofile)
	runtime.GC() // get up-to-date statistics
	runtimepprof.WriteHeapProfile(f)
	f.Close()

	profile := pprofRouter.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("", pprof.Profile)
	profile.Handle("/goroutine", pprof.Handler("goroutine"))
	profile.Handle("/threadcreate", pprof.Handler("threadcreate"))
	profile.Handle("/heap", pprof.Handler("heap"))
	profile.Handle("/block", pprof.Handler("block"))
	profile.Handle("/mutex", pprof.Handler("mutex"))

	return logMiddleware(router)
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "{\"status\": \"OK\"}")
}

func readyCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "{\"host\": \"%v\"}", r.Host)
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
