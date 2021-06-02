package app

import (
	"github.com/bearname/url-short/pkg/short/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	expectedId := domain.UrlID{}
	expectedAlias := "alias"
	expectedOriginalUrl := "https://github.com"
	expectedCreationDate := time.Now()
	url := domain.Url{
		Id:           expectedId,
		OriginalUrl:  expectedOriginalUrl,
		CreationDate: expectedCreationDate,
		Alias:        expectedAlias,
	}

	assert.Equal(t, expectedId, url.Id, "url id should be equal")
	assert.Equal(t, expectedOriginalUrl, url.OriginalUrl, "url OriginalUrl should be equal")
	assert.Equal(t, expectedCreationDate, url.CreationDate, "url CreationDate should be equal")
	assert.Equal(t, expectedAlias, url.Alias, "url Alias should be equal")
	expectedUrlString := url.Id.String() + "," + url.OriginalUrl + "," + url.CreationDate.String() + "," + url.Alias
	assert.Equal(t, expectedUrlString, url.String(), "url String should be equal")
	assert.Equal(t, expectedId.ID(), url.Id.ID(), "url id should be equal")
}
