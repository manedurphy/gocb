// Code generated by mockery v2.26.1. DO NOT EDIT.

package gocb

import mock "github.com/stretchr/testify/mock"

// mockQueryProvider is an autogenerated mock type for the queryProvider type
type mockQueryProvider struct {
	mock.Mock
}

// Query provides a mock function with given fields: statement, s, opts
func (_m *mockQueryProvider) Query(statement string, s *Scope, opts *QueryOptions) (*QueryResult, error) {
	ret := _m.Called(statement, s, opts)

	var r0 *QueryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *Scope, *QueryOptions) (*QueryResult, error)); ok {
		return rf(statement, s, opts)
	}
	if rf, ok := ret.Get(0).(func(string, *Scope, *QueryOptions) *QueryResult); ok {
		r0 = rf(statement, s, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QueryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *Scope, *QueryOptions) error); ok {
		r1 = rf(statement, s, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockQueryProvider interface {
	mock.TestingT
	Cleanup(func())
}

// newMockQueryProvider creates a new instance of mockQueryProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockQueryProvider(t mockConstructorTestingTnewMockQueryProvider) *mockQueryProvider {
	mock := &mockQueryProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
