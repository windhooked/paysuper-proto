// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	recurringpb "github.com/paysuper/paysuper-proto/go/recurringpb"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryHandler is an autogenerated mock type for the RepositoryHandler type
type RepositoryHandler struct {
	mock.Mock
}

// DeleteSavedCard provides a mock function with given fields: _a0, _a1, _a2
func (_m *RepositoryHandler) DeleteSavedCard(_a0 context.Context, _a1 *recurringpb.DeleteSavedCardRequest, _a2 *recurringpb.DeleteSavedCardResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.DeleteSavedCardRequest, *recurringpb.DeleteSavedCardResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindSavedCardById provides a mock function with given fields: _a0, _a1, _a2
func (_m *RepositoryHandler) FindSavedCardById(_a0 context.Context, _a1 *recurringpb.FindByStringValue, _a2 *recurringpb.SavedCard) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.FindByStringValue, *recurringpb.SavedCard) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindSavedCards provides a mock function with given fields: _a0, _a1, _a2
func (_m *RepositoryHandler) FindSavedCards(_a0 context.Context, _a1 *recurringpb.SavedCardRequest, _a2 *recurringpb.SavedCardList) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.SavedCardRequest, *recurringpb.SavedCardList) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertSavedCard provides a mock function with given fields: _a0, _a1, _a2
func (_m *RepositoryHandler) InsertSavedCard(_a0 context.Context, _a1 *recurringpb.SavedCardRequest, _a2 *recurringpb.Result) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *recurringpb.SavedCardRequest, *recurringpb.Result) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
