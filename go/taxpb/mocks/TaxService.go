// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/micro/go-micro/client"

	mock "github.com/stretchr/testify/mock"

	taxpb "github.com/paysuper/paysuper-proto/go/taxpb"
)

// TaxService is an autogenerated mock type for the TaxService type
type TaxService struct {
	mock.Mock
}

// CreateOrUpdate provides a mock function with given fields: ctx, in, opts
func (_m *TaxService) CreateOrUpdate(ctx context.Context, in *taxpb.TaxRate, opts ...client.CallOption) (*taxpb.TaxRate, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *taxpb.TaxRate
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.TaxRate, ...client.CallOption) *taxpb.TaxRate); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxpb.TaxRate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taxpb.TaxRate, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRateById provides a mock function with given fields: ctx, in, opts
func (_m *TaxService) DeleteRateById(ctx context.Context, in *taxpb.DeleteRateRequest, opts ...client.CallOption) (*taxpb.DeleteRateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *taxpb.DeleteRateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.DeleteRateRequest, ...client.CallOption) *taxpb.DeleteRateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxpb.DeleteRateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taxpb.DeleteRateRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRate provides a mock function with given fields: ctx, in, opts
func (_m *TaxService) GetRate(ctx context.Context, in *taxpb.GeoIdentity, opts ...client.CallOption) (*taxpb.TaxRate, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *taxpb.TaxRate
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.GeoIdentity, ...client.CallOption) *taxpb.TaxRate); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxpb.TaxRate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taxpb.GeoIdentity, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRates provides a mock function with given fields: ctx, in, opts
func (_m *TaxService) GetRates(ctx context.Context, in *taxpb.GetRatesRequest, opts ...client.CallOption) (*taxpb.GetRatesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *taxpb.GetRatesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *taxpb.GetRatesRequest, ...client.CallOption) *taxpb.GetRatesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxpb.GetRatesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *taxpb.GetRatesRequest, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
