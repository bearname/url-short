package app

type Service interface {
	CreateUrl(originalUrl string, customAlias string) error
	ReadUrl(shortUrl string) (*Url, error)
}