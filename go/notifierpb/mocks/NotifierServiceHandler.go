// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	notifierpb "github.com/paysuper/paysuper-proto/go/notifierpb"
	mock "github.com/stretchr/testify/mock"
)

// NotifierServiceHandler is an autogenerated mock type for the NotifierServiceHandler type
type NotifierServiceHandler struct {
	mock.Mock
}

// CheckUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *NotifierServiceHandler) CheckUser(_a0 context.Context, _a1 *notifierpb.CheckUserRequest, _a2 *notifierpb.CheckUserResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *notifierpb.CheckUserRequest, *notifierpb.CheckUserResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
