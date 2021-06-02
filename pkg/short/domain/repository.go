package domain

type UrlRepository interface {
	NextID() UrlID
	Create(item Url) error
	FindByAlias(alias string) (*Url, error)
}
