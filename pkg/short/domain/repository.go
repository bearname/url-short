package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UrlRepository interface {
	Create(id primitive.ObjectID, originalUrl string, customAlias string) (int64, error)
	Read(shortUrl string) (*Url, error)
}
