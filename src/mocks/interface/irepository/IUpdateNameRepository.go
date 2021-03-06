// Code generated by mockery v2.9.2. DO NOT EDIT.

package mocks

import (
	entity "TTD-golang-gin-test/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IUpdateNameRepository is an autogenerated mock type for the IUpdateNameRepository type
type IUpdateNameRepository struct {
	mock.Mock
}

// UpdateName provides a mock function with given fields: ctx, request
func (_m *IUpdateNameRepository) UpdateName(ctx context.Context, request entity.Shop) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Shop) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
