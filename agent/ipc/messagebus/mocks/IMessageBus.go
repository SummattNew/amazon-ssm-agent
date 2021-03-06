// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IMessageBus is an autogenerated mock type for the IMessageBus type
type IMessageBus struct {
	mock.Mock
}

// ProcessHealthRequest provides a mock function with given fields:
func (_m *IMessageBus) ProcessHealthRequest() {
	_m.Called()
}

// ProcessTerminationRequest provides a mock function with given fields:
func (_m *IMessageBus) ProcessTerminationRequest() {
	_m.Called()
}

// RebootRequestChannel provides a mock function with given fields:
func (_m *IMessageBus) RebootRequestChannel() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}
