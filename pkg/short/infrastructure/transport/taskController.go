package transport

import (
	"github.com/berarname/url-short/pkg/shortener/app"
)

type TaskController struct {
	Service app.Service
	BaseController
}
