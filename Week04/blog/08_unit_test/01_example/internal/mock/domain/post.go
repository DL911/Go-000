// Code generated by MockGen. DO NOT EDIT.
// Source: ./01_example/internal/domain/post.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/domain"
	reflect "reflect"
)

// MockIPostRepo is a mock of IPostRepo interface
type MockIPostRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIPostRepoMockRecorder
}

// MockIPostRepoMockRecorder is the mock recorder for MockIPostRepo
type MockIPostRepoMockRecorder struct {
	mock *MockIPostRepo
}

// NewMockIPostRepo creates a new mock instance
func NewMockIPostRepo(ctrl *gomock.Controller) *MockIPostRepo {
	mock := &MockIPostRepo{ctrl: ctrl}
	mock.recorder = &MockIPostRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPostRepo) EXPECT() *MockIPostRepoMockRecorder {
	return m.recorder
}

// MockIPostUsecase is a mock of IPostUsecase interface
type MockIPostUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIPostUsecaseMockRecorder
}

// MockIPostUsecaseMockRecorder is the mock recorder for MockIPostUsecase
type MockIPostUsecaseMockRecorder struct {
	mock *MockIPostUsecase
}

// NewMockIPostUsecase creates a new mock instance
func NewMockIPostUsecase(ctrl *gomock.Controller) *MockIPostUsecase {
	mock := &MockIPostUsecase{ctrl: ctrl}
	mock.recorder = &MockIPostUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPostUsecase) EXPECT() *MockIPostUsecaseMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method
func (m *MockIPostUsecase) CreateArticle(ctx context.Context, article domain.Article) (domain.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", ctx, article)
	ret0, _ := ret[0].(domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle
func (mr *MockIPostUsecaseMockRecorder) CreateArticle(ctx, article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockIPostUsecase)(nil).CreateArticle), ctx, article)
}
