package domain

type UrlRepository interface {
	Create(originalUrl string, customAlias string) (int64, error)
	Read(shortUrl string) (*Url, error)
}
