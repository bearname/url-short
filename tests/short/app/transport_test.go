package app

import (
	"github.com/bearname/url-short/internal/short/app"
	"github.com/bearname/url-short/internal/short/infrastructure/transport"
	"github.com/bearname/url-short/tests/short/app/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUrlController_FindByUrl_ErrUrlNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUrlService := mocks.NewMockService(mockCtrl)

	initialCreateUrlRequest := transport.NewCreateUrlRequest("bitl", "https://github.com")
	mockUrlService.EXPECT().CreateShortUrl(initialCreateUrlRequest).Return(nil, app.ErrUrlNotFound)
	//
	//appointmentService := transport.NewUrlController(mockUrlService)
	//urlRedirect, err := appointmentService.FindUrl(initialCreateUrlRequest)

	//assert.Equal(t, app.ErrUrlNotFound, err, "error not matching. expect '"+app.ErrUrlNotFound.Error()+"', actual '"+err.Error()+"'")
	//assert.Nil(t, urlRedirect, "invalid url redirect")
}
