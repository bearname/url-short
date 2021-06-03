package app

import (
	"errors"
	"github.com/bearname/url-short/pkg/common/uuid"
	"github.com/bearname/url-short/pkg/short/app"
	"github.com/bearname/url-short/pkg/short/domain"
	"github.com/bearname/url-short/pkg/short/infrastructure/transport"
	"github.com/bearname/url-short/tests/short/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlService_CreateShortUrl_ErrInvalidUrl(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAppointmentRepo := mocks.NewMockUrlRepository(mockCtrl)
	parameter := transport.CreateUrlRequest{OriginalUrl: "http/gith", CustomAlias: ""}
	appointmentService := app.NewUrlService(mockAppointmentRepo)

	shortUrl, err := appointmentService.CreateShortUrl(&parameter)
	assert.Equal(t, 0, len(shortUrl), "invalid actualShortUrl")
	if err == nil {
		t.Error("Invalid error")
		return
	}
	assert.Equal(t, app.ErrInvalidUrl, err, "error not matching. expect app.ErrInvalidUrl, actual "+err.Error())
}

func TestUrlService_CreateShortUrl_ErrDuplicateUrl(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAppointmentRepo := mocks.NewMockUrlRepository(mockCtrl)
	id := domain.UrlID{}
	mockAppointmentRepo.EXPECT().NextID().Return(id)
	request := transport.CreateUrlRequest{OriginalUrl: "https://github.com", CustomAlias: ""}

	url := domain.Url{
		Id:          id,
		OriginalUrl: request.GetOriginalUrl(),
		Alias:       "",
	}

	mockAppointmentRepo.EXPECT().Create(url).Return(app.ErrDuplicateUrl)

	appointmentService := app.NewUrlService(mockAppointmentRepo)
	shortUrl, err := appointmentService.CreateShortUrl(&request)

	assert.Equal(t, 0, len(shortUrl), "invalid actualShortUrl")
	if err != nil {
		t.Error("Invalid error")
		return
	}
	assert.Equal(t, app.ErrDuplicateUrl, err, "error not matching. expect '"+app.ErrDuplicateUrl.Error()+"', actual '"+err.Error()+"'")
}

func TestUrlService_CreateShortUrl_DatabaseError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAppointmentRepo := mocks.NewMockUrlRepository(mockCtrl)
	id := domain.UrlID{}
	mockAppointmentRepo.EXPECT().NextID().Return(id)
	request := transport.CreateUrlRequest{OriginalUrl: "https://github.com", CustomAlias: "bitl"}

	url := domain.Url{
		Id:          id,
		OriginalUrl: request.GetOriginalUrl(),
		Alias:       "bitl",
	}

	initialError := errors.New("simple protocol queries must be run with client_encoding=UTF8")
	mockAppointmentRepo.EXPECT().Create(url).Return(initialError)

	appointmentService := app.NewUrlService(mockAppointmentRepo)
	shortUrl, err := appointmentService.CreateShortUrl(&request)

	assert.Equal(t, 0, len(shortUrl), "invalid actualShortUrl")
	if err == nil {
		t.Error("Invalid error")
		return
	}
	assert.NotEqual(t, initialError.Error(), err, "error not matching. expect '"+app.ErrDuplicateUrl.Error()+"', actual '"+err.Error()+"'")
}

func TestUrlService_CreateShortUrl_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAppointmentRepo := mocks.NewMockUrlRepository(mockCtrl)
	id := domain.UrlID(uuid.Generate())
	mockAppointmentRepo.EXPECT().NextID().Return(id)
	initialAlias := "bitl"
	request := transport.CreateUrlRequest{OriginalUrl: "https://github.com", CustomAlias: initialAlias}

	url := domain.Url{
		Id:          id,
		OriginalUrl: request.GetOriginalUrl(),
		Alias:       initialAlias,
	}

	mockAppointmentRepo.EXPECT().Create(url).Return(nil)

	appointmentService := app.NewUrlService(mockAppointmentRepo)
	shortUrl, err := appointmentService.CreateShortUrl(&request)

	assert.Equal(t, initialAlias, shortUrl, "invalid actualShortUrl")
	if err != nil {
		t.Error("error not matching. not expect any error, but actual '"+err.Error()+"'", err)
	}
}

func TestUrlService_FindByUrl_ErrUrlNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAppointmentRepo := mocks.NewMockUrlRepository(mockCtrl)

	shortUrl := "bitl"
	mockAppointmentRepo.EXPECT().FindByAlias(shortUrl).Return(nil, app.ErrUrlNotFound)

	appointmentService := app.NewUrlService(mockAppointmentRepo)
	urlRedirect, err := appointmentService.FindUrl(shortUrl)

	if err != nil {
		t.Error("Invalid error")
		return
	}
	assert.Equal(t, app.ErrUrlNotFound, err, "error not matching. expect '"+app.ErrUrlNotFound.Error()+"', actual '"+err.Error()+"'")
	assert.Nil(t, urlRedirect, "invalid url redirect")
}

func TestUrlService_FindByUrl_Valid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAppointmentRepo := mocks.NewMockUrlRepository(mockCtrl)

	shortUrl := "bitl"
	id := domain.UrlID(uuid.Generate())
	expectedUrl := domain.Url{
		Id:          id,
		OriginalUrl: "https://github.com",
		Alias:       shortUrl,
	}

	mockAppointmentRepo.EXPECT().FindByAlias(shortUrl).Return(&expectedUrl, nil)

	appointmentService := app.NewUrlService(mockAppointmentRepo)
	actualUrl, err := appointmentService.FindUrl(shortUrl)

	if actualUrl == nil {
		t.Error("actualUrl must be not nil")
		return
	}
	assert.Equal(t, expectedUrl.OriginalUrl, actualUrl.OriginalUrl, "original url not match")
	assert.Equal(t, expectedUrl.Alias, actualUrl.Alias, "alias url not match")
	if err != nil {
		t.Error("error must be nil, but "+err.Error(), err)
	}
}
