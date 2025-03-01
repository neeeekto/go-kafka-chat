// Code generated by mockery v2.49.1. DO NOT EDIT.

package contracts

import (
	context "context"
	entities "go-kafka-chat/intenral/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockUsersRepository is an autogenerated mock type for the UsersRepository type
type MockUsersRepository struct {
	mock.Mock
}

// CheckExists provides a mock function with given fields: ctx, userID
func (_m *MockUsersRepository) CheckExists(ctx context.Context, userID ...entities.UserID) (bool, error) {
	_va := make([]interface{}, len(userID))
	for _i := range userID {
		_va[_i] = userID[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CheckExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...entities.UserID) (bool, error)); ok {
		return rf(ctx, userID...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...entities.UserID) bool); ok {
		r0 = rf(ctx, userID...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...entities.UserID) error); ok {
		r1 = rf(ctx, userID...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, user
func (_m *MockUsersRepository) Save(ctx context.Context, user *entities.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUsersRepository creates a new instance of MockUsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUsersRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUsersRepository {
	mock := &MockUsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
