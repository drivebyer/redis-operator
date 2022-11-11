// Code generated by mockery v2.9.6. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	redisfailoverv1 "github.com/spotahome/redis-operator/api/redisfailover/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// RedisFailover is an autogenerated mock type for the RedisFailover type
type RedisFailover struct {
	mock.Mock
}

// ListRedisFailovers provides a mock function with given fields: ctx, namespace, opts
func (_m *RedisFailover) ListRedisFailovers(ctx context.Context, namespace string, opts v1.ListOptions) (*redisfailoverv1.RedisFailoverList, error) {
	ret := _m.Called(ctx, namespace, opts)

	var r0 *redisfailoverv1.RedisFailoverList
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.ListOptions) *redisfailoverv1.RedisFailoverList); ok {
		r0 = rf(ctx, namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redisfailoverv1.RedisFailoverList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.ListOptions) error); ok {
		r1 = rf(ctx, namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRedisFailovers provides a mock function with given fields: ctx, namespace, opts
func (_m *RedisFailover) WatchRedisFailovers(ctx context.Context, namespace string, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, namespace, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.ListOptions) error); ok {
		r1 = rf(ctx, namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRedisFailover interface {
	mock.TestingT
	Cleanup(func())
}

// NewRedisFailover creates a new instance of RedisFailover. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRedisFailover(t mockConstructorTestingTNewRedisFailover) *RedisFailover {
	mock := &RedisFailover{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
