package app

import (
	"errors"
	"github.com/bearname/url-short/pkg/short/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = int64(len(alphabet))
)

var (
	ErrUrlNotFound  = errors.New("url not found")
	ErrDuplicateUrl = errors.New("url with such SKU already exists")
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
	id := primitive.NewObjectID()
	short := s.encode(int64(id.Timestamp().Second()))
	if len(customAlias) == 0 {
		customAlias = short
	}
	_, err := s.repo.Create(id, originalUrl, customAlias)
	if err != nil {
		return "", nil
	}
	return short, nil
}

func (s *UrlService) ReadUrl(shortUrl string) (*domain.Url, error) {
	//decode, err := s.decode(shortUrl)
	//if err != nil {
	//	return nil, err
	//}
	//
	//unixTimeUTC := time.Unix(decode, 0)
	//
	//objectId := primitive.NewObjectIDFromTimestamp(unixTimeUTC)
	//s2 := objectId.String()
	//fmt.Println(s2)

	return s.repo.Read(shortUrl)
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
