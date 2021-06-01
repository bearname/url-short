package transport

import (
	"github.com/bearname/url-short/pkg/shortener/app"
)

type TaskController struct {
	Service app.Service
	BaseController
}
