package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type UrlID uuid.UUID

func (u UrlID) String() string {
	return uuid.UUID(u).String()
}

type Url struct {
	Id             UrlID
	OriginalUrl    string
	CreationDate   time.Time
	ExpirationDate time.Time
	CustomUrl      string
}

func NewUrl(originalUrl string, creationDate time.Time, expirationDate time.Time) *Url {
	u := new(Url)
	u.OriginalUrl = originalUrl
	u.CreationDate = creationDate
	u.ExpirationDate = expirationDate
	return u
}
