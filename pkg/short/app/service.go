package app

import (
	"github.com/bearname/url-short/pkg/short/domain"
	"github.com/bearname/url-short/pkg/short/infrastructure/util"
	"github.com/pkg/errors"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint32(len(alphabet))
)

var (
	ErrUrlNotFound  = errors.New("url not found")
	ErrDuplicateUrl = errors.New("url with such OriginalUrl already exists")
	ErrInvalidUrl   = errors.New("url not valid")
)

type UrlService struct {
	repo domain.UrlRepository
}

func NewUrlService(repo domain.UrlRepository) *UrlService {
	u := new(UrlService)
	u.repo = repo
	return u
}

func (s *UrlService) CreateShortUrl(parameter domain.UrlParameter) (string, error) {
	id := s.repo.NextID()

	isValid := util.IsValidUrl(parameter.GetOriginalUrl())
	if !isValid {
		return "", ErrInvalidUrl
	}

	item := s.buildUrlItem(parameter, id)

	err := s.repo.Create(item)
	if err != nil {
		if err.Error() == "url with such OriginalUrl already exists" {
			return "", ErrDuplicateUrl
		}
		return "", errors.WithStack(err)
	}

	return item.Alias, nil
}

func (s *UrlService) FindUrl(shortUrl string) (*domain.Url, error) {
	id, err := s.repo.FindByAlias(shortUrl)
	if err != nil {
		return nil, ErrUrlNotFound
	}

	return id, nil
}

func (s *UrlService) buildUrlItem(parameter domain.UrlParameter, id domain.UrlID) domain.Url {
	item := domain.Url{
		Id:          id,
		OriginalUrl: parameter.GetOriginalUrl(),
		Alias:       parameter.GetCustomAlias(),
	}

	if len(parameter.GetCustomAlias()) == 0 {
		shortAlias := s.createShortUrl(id.ID())
		if len(item.Alias) == 0 {
			item.Alias = shortAlias
		}
	}

	return item
}

func (s *UrlService) createShortUrl(number uint32) string {
	var encodedBuilder strings.Builder

	for ; number > 0; number = number / length {
		encodedBuilder.WriteString(string(alphabet[(number % length)]))
	}

	return encodedBuilder.String()
}
