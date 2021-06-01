package transport

import (
	"github.com/bearname/url-short/pkg/short/app"
	"github.com/bearname/url-short/pkg/short/domain"
	"github.com/bearname/url-short/pkg/short/infrastructure/middleware"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UrlController struct {
	service domain.Service
	BaseController
}

func NewUrlController(service *app.UrlService) *UrlController {
	u := new(UrlController)
	u.service = service
	return u
}

func (c *UrlController) Create() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		url, ok := context.Get(request, "url").(middleware.UrlRequest)
		if !ok {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("url not present"))
			return
		}

		shortUrl, err := c.service.CreateUrl(url.Url, url.CustomAlias)
		if err != nil {
			log.Error(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("failed create short url"))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte("{\"short_url\": \"http://localhost:8000/" + shortUrl + "\"}"))
		return
	}
}

func (c *UrlController) Redirect() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		shortUrl, ok := vars["shortUrl"]
		if !ok {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("url not present"))
			return
		}

		url, err := c.service.ReadUrl(shortUrl)
		if err != nil {
			log.Error(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("failed create short url"))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Location", url.OriginalUrl)
		return
	}
}
