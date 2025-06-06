// Code generated by MockGen. DO NOT EDIT.
// Source: cart_repository.go
//
// Generated by this command:
//
//	mockgen -source=cart_repository.go -destination=./mocks/cart_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	repository "github.com/dbiagi/shopping-bag/internal/cart/repository"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockCartRepositoryInterface is a mock of CartRepositoryInterface interface.
type MockCartRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockCartRepositoryInterfaceMockRecorder is the mock recorder for MockCartRepositoryInterface.
type MockCartRepositoryInterfaceMockRecorder struct {
	mock *MockCartRepositoryInterface
}

// NewMockCartRepositoryInterface creates a new mock instance.
func NewMockCartRepositoryInterface(ctrl *gomock.Controller) *MockCartRepositoryInterface {
	mock := &MockCartRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockCartRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartRepositoryInterface) EXPECT() *MockCartRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CartByID mocks base method.
func (m *MockCartRepositoryInterface) CartByID(id uuid.UUID) (*repository.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CartByID", id)
	ret0, _ := ret[0].(*repository.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CartByID indicates an expected call of CartByID.
func (mr *MockCartRepositoryInterfaceMockRecorder) CartByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CartByID", reflect.TypeOf((*MockCartRepositoryInterface)(nil).CartByID), id)
}
