// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

type Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *Interface) EXPECT() *Interface_Expecter {
	return &Interface_Expecter{mock: &_m.Mock}
}

// GetValue provides a mock function with given fields: _a0, _a1, _a2
func (_m *Interface) GetValue(_a0 int, _a1 string, _a2 time.Duration) (string, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetValue")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string, time.Duration) (string, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(int, string, time.Duration) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int, string, time.Duration) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interface_GetValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValue'
type Interface_GetValue_Call struct {
	*mock.Call
}

// GetValue is a helper method to define mock.On call
//   - _a0 int
//   - _a1 string
//   - _a2 time.Duration
func (_e *Interface_Expecter) GetValue(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Interface_GetValue_Call {
	return &Interface_GetValue_Call{Call: _e.mock.On("GetValue", _a0, _a1, _a2)}
}

func (_c *Interface_GetValue_Call) Run(run func(_a0 int, _a1 string, _a2 time.Duration)) *Interface_GetValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string), args[2].(time.Duration))
	})
	return _c
}

func (_c *Interface_GetValue_Call) Return(_a0 string, _a1 error) *Interface_GetValue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Interface_GetValue_Call) RunAndReturn(run func(int, string, time.Duration) (string, error)) *Interface_GetValue_Call {
	_c.Call.Return(run)
	return _c
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}