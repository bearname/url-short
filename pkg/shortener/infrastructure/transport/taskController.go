package transport

import (
	"github.com/berarname/url-shortener/pkg/shortener/app"
)

type TaskController struct {
	Service app.Service
	BaseController
}
