package domain

type UrlParameter interface {
	GetCustomAlias() string
	GetOriginalUrl() string
}

type Service interface {
	CreateShortUrl(parameter UrlParameter) (string, error)
	FindUrl(shortUrl string) (*Url, error)
}
