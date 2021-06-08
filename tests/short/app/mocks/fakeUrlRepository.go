package mocks

import (
	domain "github.com/bearname/url-short/internal/short/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUrlRepository is a mock of UrlRepository interface.
type MockUrlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUrlRepositoryMockRecorder
}

// MockUrlRepositoryMockRecorder is the mock recorder for MockUrlRepository.
type MockUrlRepositoryMockRecorder struct {
	mock *MockUrlRepository
}

// NewMockUrlRepository creates a new mock instance.
func NewMockUrlRepository(ctrl *gomock.Controller) *MockUrlRepository {
	mock := &MockUrlRepository{ctrl: ctrl}
	mock.recorder = &MockUrlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlRepository) EXPECT() *MockUrlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUrlRepository) Create(item domain.Url) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUrlRepositoryMockRecorder) Create(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUrlRepository)(nil).Create), item)
}

// FindByAlias mocks base method.
func (m *MockUrlRepository) FindByAlias(alias string) (*domain.Url, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAlias", alias)
	ret0, _ := ret[0].(*domain.Url)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAlias indicates an expected call of FindByAlias.
func (mr *MockUrlRepositoryMockRecorder) FindByAlias(alias interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAlias", reflect.TypeOf((*MockUrlRepository)(nil).FindByAlias), alias)
}

// NextID mocks base method.
func (m *MockUrlRepository) NextID() domain.UrlID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(domain.UrlID)
	return ret0
}

// NextID indicates an expected call of NextID.
func (mr *MockUrlRepositoryMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MockUrlRepository)(nil).NextID))
}
