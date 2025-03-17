// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	resty "github.com/go-resty/resty/v2"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, user, follower
func (_m *Service) Apply(ctx context.Context, user string, follower string) (*resty.Response, error) {
	ret := _m.Called(ctx, user, follower)

	if len(ret) == 0 {
		panic("no return value specified for Apply")
	}

	var r0 *resty.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*resty.Response, error)); ok {
		return rf(ctx, user, follower)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *resty.Response); ok {
		r0 = rf(ctx, user, follower)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resty.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, user, follower)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
