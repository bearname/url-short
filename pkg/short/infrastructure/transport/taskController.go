package transport

import (
	"github.com/bearname/url-short/pkg/short/app"
)

type TaskController struct {
	Service app.Service
	BaseController
}
