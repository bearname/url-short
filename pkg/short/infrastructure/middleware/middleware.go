package middleware

import (
	"encoding/json"
	"github.com/gorilla/context"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UrlRequest struct {
	Url         string `json:"url"`
	CustomAlias string `json:"custom_alias"`
}

func DecodeCreateUrlRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if (*request).Method == "OPTIONS" {
			writer.WriteHeader(http.StatusNoContent)
			return
		}

		log.Println("DecodeCreateUrlRequest")
		var urlRequest UrlRequest
		err := json.NewDecoder(request.Body).Decode(&urlRequest)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		context.Set(request, "url", urlRequest)

		log.Println("success")

		next(writer, request)
	}
}
