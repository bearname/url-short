package app

type urlRepository interface {
	Create(originalUrl string, customAlias string) error
	Read(shortUrl string) (*Url, error)
}
