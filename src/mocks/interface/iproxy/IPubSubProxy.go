// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity2 "TTD-golang-gin-test/entity"
	mock "github.com/stretchr/testify/mock"
)

// IPubSubProxy is an autogenerated mock type for the IPubSubProxy type
type IPubSubProxy struct {
	mock.Mock
}

// PubShop provides a mock function with given fields: shop
func (_m *IPubSubProxy) PubShop(shop entity2.Shop) error {
	ret := _m.Called(shop)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity2.Shop) error); ok {
		r0 = rf(shop)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
