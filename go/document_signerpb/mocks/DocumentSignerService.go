// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import client "github.com/micro/go-micro/client"
import context "context"
import document_signerpb "github.com/paysuper/paysuper-proto/go/document_signerpb"
import mock "github.com/stretchr/testify/mock"

// DocumentSignerService is an autogenerated mock type for the DocumentSignerService type
type DocumentSignerService struct {
	mock.Mock
}

// CreateSignature provides a mock function with given fields: ctx, in, opts
func (_m *DocumentSignerService) CreateSignature(ctx context.Context, in *document_signerpb.CreateSignatureRequest, opts ...client.CallOption) (*document_signerpb.CreateSignatureResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *document_signerpb.CreateSignatureResponse
	if rf, ok := ret.Get(0).(func(context.Context, *document_signerpb.CreateSignatureRequest, ...client.CallOption) *document_signerpb.CreateSignatureResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*document_signerpb.CreateSignatureResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *document_signerpb.CreateSignatureRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignatureUrl provides a mock function with given fields: ctx, in, opts
func (_m *DocumentSignerService) GetSignatureUrl(ctx context.Context, in *document_signerpb.GetSignatureUrlRequest, opts ...client.CallOption) (*document_signerpb.GetSignatureUrlResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *document_signerpb.GetSignatureUrlResponse
	if rf, ok := ret.Get(0).(func(context.Context, *document_signerpb.GetSignatureUrlRequest, ...client.CallOption) *document_signerpb.GetSignatureUrlResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*document_signerpb.GetSignatureUrlResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *document_signerpb.GetSignatureUrlRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
