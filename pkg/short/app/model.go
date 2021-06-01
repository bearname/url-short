package app

import "time"

type Url struct {
	originalUrl    string
	creationDate   time.Time
	expirationDate time.Time
}
