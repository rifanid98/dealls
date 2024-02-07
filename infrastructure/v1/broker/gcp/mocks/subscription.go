// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	iam "cloud.google.com/go/iam"

	internalpubsub "cloud.google.com/go/internal/pubsub"

	mock "github.com/stretchr/testify/mock"

	pubsub "cloud.google.com/go/pubsub"
)

// Subscription is an autogenerated mock type for the Subscription type
type Subscription struct {
	mock.Mock
}

// Config provides a mock function with given fields: ctx
func (_m *Subscription) Config(ctx context.Context) (pubsub.SubscriptionConfig, error) {
	ret := _m.Called(ctx)

	var r0 pubsub.SubscriptionConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pubsub.SubscriptionConfig, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pubsub.SubscriptionConfig); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(pubsub.SubscriptionConfig)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx
func (_m *Subscription) Delete(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx
func (_m *Subscription) Exists(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IAM provides a mock function with given fields:
func (_m *Subscription) IAM() *iam.Handle {
	ret := _m.Called()

	var r0 *iam.Handle
	if rf, ok := ret.Get(0).(func() *iam.Handle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.Handle)
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Subscription) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Receive provides a mock function with given fields: ctx, f
func (_m *Subscription) Receive(ctx context.Context, f func(context.Context, *internalpubsub.Message)) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context, *internalpubsub.Message)) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Subscription) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, cfg
func (_m *Subscription) Update(ctx context.Context, cfg pubsub.SubscriptionConfigToUpdate) (pubsub.SubscriptionConfig, error) {
	ret := _m.Called(ctx, cfg)

	var r0 pubsub.SubscriptionConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pubsub.SubscriptionConfigToUpdate) (pubsub.SubscriptionConfig, error)); ok {
		return rf(ctx, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pubsub.SubscriptionConfigToUpdate) pubsub.SubscriptionConfig); ok {
		r0 = rf(ctx, cfg)
	} else {
		r0 = ret.Get(0).(pubsub.SubscriptionConfig)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pubsub.SubscriptionConfigToUpdate) error); ok {
		r1 = rf(ctx, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSubscription creates a new instance of Subscription. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscription(t interface {
	mock.TestingT
	Cleanup(func())
}) *Subscription {
	mock := &Subscription{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
