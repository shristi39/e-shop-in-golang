// Code generated by MockGen. DO NOT EDIT.
// Source: internal/models/productdb.go

// Package mockdb is a generated GoMock package.
package mockdb

import (
	models "commerce/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProductModels is a mock of ProductModels interface.
type MockProductModels struct {
	ctrl     *gomock.Controller
	recorder *MockProductModelsMockRecorder
}

// MockProductModelsMockRecorder is the mock recorder for MockProductModels.
type MockProductModelsMockRecorder struct {
	mock *MockProductModels
}

// NewMockProductModels creates a new mock instance.
func NewMockProductModels(ctrl *gomock.Controller) *MockProductModels {
	mock := &MockProductModels{ctrl: ctrl}
	mock.recorder = &MockProductModelsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductModels) EXPECT() *MockProductModelsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductModels) Create(name string, price int, image, description, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price, image, description, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProductModelsMockRecorder) Create(name, price, image, description, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductModels)(nil).Create), name, price, image, description, email)
}

// Delete mocks base method.
func (m *MockProductModels) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductModelsMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductModels)(nil).Delete), id)
}

// FindAllProduct mocks base method.
func (m *MockProductModels) FindAllProduct() ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllProduct")
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllProduct indicates an expected call of FindAllProduct.
func (mr *MockProductModelsMockRecorder) FindAllProduct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllProduct", reflect.TypeOf((*MockProductModels)(nil).FindAllProduct))
}

// FindProductById mocks base method.
func (m *MockProductModels) FindProductById(id int) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductById", id)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductById indicates an expected call of FindProductById.
func (mr *MockProductModelsMockRecorder) FindProductById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductById", reflect.TypeOf((*MockProductModels)(nil).FindProductById), id)
}

// FindProductByStatus mocks base method.
func (m *MockProductModels) FindProductByStatus(status bool) (*[]models.Product, *int, *int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductByStatus", status)
	ret0, _ := ret[0].(*[]models.Product)
	ret1, _ := ret[1].(*int)
	ret2, _ := ret[2].(*int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// FindProductByStatus indicates an expected call of FindProductByStatus.
func (mr *MockProductModelsMockRecorder) FindProductByStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductByStatus", reflect.TypeOf((*MockProductModels)(nil).FindProductByStatus), status)
}

// Update mocks base method.
func (m *MockProductModels) Update(status bool, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", status, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductModelsMockRecorder) Update(status, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductModels)(nil).Update), status, id)
}