// Code generated by mockery v2.26.1. DO NOT EDIT.

package gocb

import (
	context "context"

	gocbcore "github.com/couchbase/gocbcore/v10"
	mock "github.com/stretchr/testify/mock"
)

// mockSearchProviderCoreProvider is an autogenerated mock type for the searchProviderCoreProvider type
type mockSearchProviderCoreProvider struct {
	mock.Mock
}

// SearchQuery provides a mock function with given fields: ctx, opts
func (_m *mockSearchProviderCoreProvider) SearchQuery(ctx context.Context, opts gocbcore.SearchQueryOptions) (searchRowReader, error) {
	ret := _m.Called(ctx, opts)

	var r0 searchRowReader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, gocbcore.SearchQueryOptions) (searchRowReader, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, gocbcore.SearchQueryOptions) searchRowReader); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(searchRowReader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, gocbcore.SearchQueryOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockSearchProviderCoreProvider interface {
	mock.TestingT
	Cleanup(func())
}

// newMockSearchProviderCoreProvider creates a new instance of mockSearchProviderCoreProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockSearchProviderCoreProvider(t mockConstructorTestingTnewMockSearchProviderCoreProvider) *mockSearchProviderCoreProvider {
	mock := &mockSearchProviderCoreProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
