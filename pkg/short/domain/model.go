package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Url struct {
	Id             primitive.ObjectID
	OriginalUrl    string
	CreationDate   string
	ExpirationDate string
	CustomUrl      string
}

func NewUrl(originalUrl string, creationDate time.Time, expirationDate time.Time) *Url {
	u := new(Url)
	u.OriginalUrl = originalUrl
	u.CreationDate = creationDate.String()
	u.ExpirationDate = expirationDate.String()
	return u
}
