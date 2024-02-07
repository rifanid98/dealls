// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	bson "go.mongodb.org/mongo-driver/bson"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"

	mongodb "dealls/infrastructure/v1/persistence/mongodb"

	options "go.mongodb.org/mongo-driver/mongo/options"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	time "time"
)

// Session is an autogenerated mock type for the Session type
type Session struct {
	mock.Mock
}

// AbortTransaction provides a mock function with given fields: _a0
func (_m *Session) AbortTransaction(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdvanceClusterTime provides a mock function with given fields: _a0
func (_m *Session) AdvanceClusterTime(_a0 bson.Raw) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(bson.Raw) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AdvanceOperationTime provides a mock function with given fields: _a0
func (_m *Session) AdvanceOperationTime(_a0 *primitive.Timestamp) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*primitive.Timestamp) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client provides a mock function with given fields:
func (_m *Session) Client() mongodb.Client {
	ret := _m.Called()

	var r0 mongodb.Client
	if rf, ok := ret.Get(0).(func() mongodb.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongodb.Client)
		}
	}

	return r0
}

// ClusterTime provides a mock function with given fields:
func (_m *Session) ClusterTime() bson.Raw {
	ret := _m.Called()

	var r0 bson.Raw
	if rf, ok := ret.Get(0).(func() bson.Raw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	return r0
}

// CommitTransaction provides a mock function with given fields: _a0
func (_m *Session) CommitTransaction(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deadline provides a mock function with given fields:
func (_m *Session) Deadline() (time.Time, bool) {
	ret := _m.Called()

	var r0 time.Time
	var r1 bool
	if rf, ok := ret.Get(0).(func() (time.Time, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Done provides a mock function with given fields:
func (_m *Session) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// EndSession provides a mock function with given fields: _a0
func (_m *Session) EndSession(_a0 context.Context) {
	_m.Called(_a0)
}

// Err provides a mock function with given fields:
func (_m *Session) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OperationTime provides a mock function with given fields:
func (_m *Session) OperationTime() *primitive.Timestamp {
	ret := _m.Called()

	var r0 *primitive.Timestamp
	if rf, ok := ret.Get(0).(func() *primitive.Timestamp); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*primitive.Timestamp)
		}
	}

	return r0
}

// StartTransaction provides a mock function with given fields: _a0
func (_m *Session) StartTransaction(_a0 ...*options.TransactionOptions) error {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*options.TransactionOptions) error); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Value provides a mock function with given fields: key
func (_m *Session) Value(key interface{}) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// WithTransaction provides a mock function with given fields: ctx, fn, opts
func (_m *Session) WithTransaction(ctx context.Context, fn func(mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, fn)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, func(mongo.SessionContext) (interface{}, error), ...*options.TransactionOptions) (interface{}, error)); ok {
		return rf(ctx, fn, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, func(mongo.SessionContext) (interface{}, error), ...*options.TransactionOptions) interface{}); ok {
		r0 = rf(ctx, fn, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, func(mongo.SessionContext) (interface{}, error), ...*options.TransactionOptions) error); ok {
		r1 = rf(ctx, fn, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSession creates a new instance of Session. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSession(t interface {
	mock.TestingT
	Cleanup(func())
}) *Session {
	mock := &Session{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
