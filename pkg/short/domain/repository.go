package domain

type UrlRepository interface {
	NextID() UrlID
	Create(item Url) error
	FindById(id UrlID) (*Url, error)
}
