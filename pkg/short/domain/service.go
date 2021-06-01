package domain

type Service interface {
	CreateUrl(originalUrl string, customAlias string) (string, error)
	ReadUrl(shortUrl string) (*Url, error)
}
