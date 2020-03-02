// Code generated by mockery v1.0.0. DO NOT EDIT.

package gocb

import gocbcore "github.com/couchbase/gocbcore/v8"
import mock "github.com/stretchr/testify/mock"

// mockQueryProvider is an autogenerated mock type for the queryProvider type
type mockQueryProvider struct {
	mock.Mock
}

// N1QLQuery provides a mock function with given fields: opts
func (_m *mockQueryProvider) N1QLQuery(opts gocbcore.N1QLQueryOptions) (queryRowReader, error) {
	ret := _m.Called(opts)

	var r0 queryRowReader
	if rf, ok := ret.Get(0).(func(gocbcore.N1QLQueryOptions) queryRowReader); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queryRowReader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gocbcore.N1QLQueryOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
