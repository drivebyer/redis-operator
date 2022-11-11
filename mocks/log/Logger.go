// Code generated by mockery v2.9.6. DO NOT EDIT.

package mocks

import (
	log "github.com/spotahome/redis-operator/log"
	mock "github.com/stretchr/testify/mock"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: _a0
func (_m *Logger) Debug(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: _a0, _a1
func (_m *Logger) Debugf(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// Debugln provides a mock function with given fields: _a0
func (_m *Logger) Debugln(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: _a0
func (_m *Logger) Error(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: _a0, _a1
func (_m *Logger) Errorf(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// Errorln provides a mock function with given fields: _a0
func (_m *Logger) Errorln(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: _a0
func (_m *Logger) Fatal(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: _a0, _a1
func (_m *Logger) Fatalf(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// Fatalln provides a mock function with given fields: _a0
func (_m *Logger) Fatalln(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: _a0
func (_m *Logger) Info(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: _a0, _a1
func (_m *Logger) Infof(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// Infoln provides a mock function with given fields: _a0
func (_m *Logger) Infoln(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Panic provides a mock function with given fields: _a0
func (_m *Logger) Panic(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Panicf provides a mock function with given fields: _a0, _a1
func (_m *Logger) Panicf(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// Panicln provides a mock function with given fields: _a0
func (_m *Logger) Panicln(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Set provides a mock function with given fields: level
func (_m *Logger) Set(level log.Level) error {
	ret := _m.Called(level)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.Level) error); ok {
		r0 = rf(level)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Warn provides a mock function with given fields: _a0
func (_m *Logger) Warn(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// Warnf provides a mock function with given fields: _a0, _a1
func (_m *Logger) Warnf(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// Warningf provides a mock function with given fields: _a0, _a1
func (_m *Logger) Warningf(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// Warnln provides a mock function with given fields: _a0
func (_m *Logger) Warnln(_a0 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	_m.Called(_ca...)
}

// With provides a mock function with given fields: key, value
func (_m *Logger) With(key string, value interface{}) log.Logger {
	ret := _m.Called(key, value)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(string, interface{}) log.Logger); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// WithField provides a mock function with given fields: key, value
func (_m *Logger) WithField(key string, value interface{}) log.Logger {
	ret := _m.Called(key, value)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(string, interface{}) log.Logger); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// WithFields provides a mock function with given fields: values
func (_m *Logger) WithFields(values map[string]interface{}) log.Logger {
	ret := _m.Called(values)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(map[string]interface{}) log.Logger); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

type mockConstructorTestingTNewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTNewLogger) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
