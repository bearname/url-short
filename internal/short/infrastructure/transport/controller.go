package transport

import (
	"encoding/json"
	"github.com/bearname/url-short/internal/short/app"
	"github.com/bearname/url-short/internal/short/domain"
	"github.com/bearname/url-short/internal/short/infrastructure/util"
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
		log.Println("DecodeCreateUrlRequest")
		var urlRequest CreateUrlRequest
		err := json.NewDecoder(request.Body).Decode(&urlRequest)
		if err != nil {
			http.Error(writer, "customAlias or originalUrl not present on json", http.StatusBadRequest)
			return
		}

		b := !util.IsValidUrl(urlRequest.GetOriginalUrl())
		if b {
			http.Error(writer, "invalid original url", http.StatusBadRequest)
			return
		}

		//context.Set(request, "url", urlRequest)
		//url, ok := context.Get(request, "url").(CreateUrlRequest)
		//if !ok {
		//	c.BaseController.WriteError(writer, ErrBadRequest)
		//	return
		//}

		shortUrl, err := c.service.CreateShortUrl(&urlRequest)
		if err != nil {
			c.BaseController.WriteError(writer, err)
			return
		}

		c.BaseController.WriteJsonResponse(writer, CreateUrlResponse{"http://" + request.Host + "/api/v1/url/" + shortUrl})
	}
}

func (c *UrlController) Redirect() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		shortUrl, ok := vars["shortUrl"]
		if !ok {
			c.BaseController.WriteError(writer, ErrBadRequest)
			return
		}

		url, err := c.service.FindUrl(shortUrl)
		if err != nil {
			c.BaseController.WriteError(writer, err)
			return
		}

		http.Redirect(writer, request, url.OriginalUrl, http.StatusPermanentRedirect)
	}
}
