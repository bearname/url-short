package domain

import (
	"github.com/google/uuid"
	"time"
)

type UrlID uuid.UUID

func (u UrlID) String() string {
	return uuid.UUID(u).String()
}

func (u UrlID) ID() uint32 {
	return uuid.UUID(u).ID()
}

type Url struct {
	Id           UrlID
	OriginalUrl  string
	CreationDate time.Time
	Alias        string
}

func (u *Url) String() string {
	return u.Id.String() + "," + u.OriginalUrl + "," + u.CreationDate.String() + "," + u.Alias
}
