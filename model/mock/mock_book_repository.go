// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rhtyx/bayarind-service.git/model (interfaces: BookRepository)
//
// Generated by this command:
//
//	mockgen -destination=model/mock/mock_book_repository.go -package=mock github.com/rhtyx/bayarind-service.git/model BookRepository
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/rhtyx/bayarind-service.git/model"
	gomock "go.uber.org/mock/gomock"
)

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBookRepository) Create(arg0 context.Context, arg1 *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBookRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockBookRepository) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBookRepositoryMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookRepository)(nil).Delete), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockBookRepository) FindAll(arg0 context.Context) ([]*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockBookRepositoryMockRecorder) FindAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockBookRepository)(nil).FindAll), arg0)
}

// FindByID mocks base method.
func (m *MockBookRepository) FindByID(arg0 context.Context, arg1 int64) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockBookRepositoryMockRecorder) FindByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockBookRepository)(nil).FindByID), arg0, arg1)
}

// FindByISBN mocks base method.
func (m *MockBookRepository) FindByISBN(arg0 context.Context, arg1 string) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByISBN", arg0, arg1)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByISBN indicates an expected call of FindByISBN.
func (mr *MockBookRepositoryMockRecorder) FindByISBN(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByISBN", reflect.TypeOf((*MockBookRepository)(nil).FindByISBN), arg0, arg1)
}

// Update mocks base method.
func (m *MockBookRepository) Update(arg0 context.Context, arg1 *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBookRepositoryMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookRepository)(nil).Update), arg0, arg1)
}
