package app

import (
	"errors"
	"fmt"
	"github.com/bearname/url-short/pkg/short/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = int64(len(alphabet))
)

type UrlService struct {
	repo domain.UrlRepository
}

func NewUrlService(repo domain.UrlRepository) *UrlService {
	u := new(UrlService)
	u.repo = repo
	return u
}

func (s *UrlService) CreateUrl(originalUrl string, customAlias string) (string, error) {
	i, err := s.repo.Create(originalUrl, customAlias)
	if len(customAlias) != 0 {
		return customAlias, err
	}
	short := s.encode(i)
	return short, nil
}

func (s *UrlService) ReadUrl(shortUrl string) (*domain.Url, error) {
	decode, err := s.decode(shortUrl)
	if err != nil {
		return nil, err
	}

	unix := time.Unix(decode, 0)
	objectId := primitive.NewObjectIDFromTimestamp(unix)
	s2 := objectId.String()
	fmt.Println(s2)
	return s.repo.Read(s2)
}

func (s *UrlService) encode(number int64) string {
	var encodedBuilder strings.Builder

	for ; number > 0; number = number / length {
		encodedBuilder.WriteString(string(alphabet[(number % length)]))
	}

	return encodedBuilder.String()
}

func (s *UrlService) decode(encoded string) (int64, error) {
	var number int64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return int64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += int64(alphabeticPosition) * int64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
