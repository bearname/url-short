package transport

import (
	"github.com/bearname/url-short/pkg/short/app"
	"github.com/bearname/url-short/pkg/short/domain"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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
		url, ok := context.Get(request, "url").(CreateUrlRequest)
		if !ok {
			c.BaseController.WriteError(writer, ErrBadRequest)
			return
		}

		shortUrl, err := c.service.CreateShortUrl(&url)
		if err != nil {
			c.BaseController.WriteError(writer, err)
			return
		}

		c.WriteJsonResponse(writer, CreateUrlResponse{"http://" + request.Host + "/api/v1/url/" + shortUrl})
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

		http.Redirect(writer, request, url.OriginalUrl, http.StatusTemporaryRedirect)
	}
}
