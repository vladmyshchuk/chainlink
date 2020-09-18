// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"
	mock "github.com/stretchr/testify/mock"
)

// Task is an autogenerated mock type for the Task type
type Task struct {
	mock.Mock
}

// OutputTask provides a mock function with given fields:
func (_m *Task) OutputTask() pipeline.Task {
	ret := _m.Called()

	var r0 pipeline.Task
	if rf, ok := ret.Get(0).(func() pipeline.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.Task)
		}
	}

	return r0
}

// Run provides a mock function with given fields: inputs
func (_m *Task) Run(inputs []pipeline.Result) pipeline.Result {
	ret := _m.Called(inputs)

	var r0 pipeline.Result
	if rf, ok := ret.Get(0).(func([]pipeline.Result) pipeline.Result); ok {
		r0 = rf(inputs)
	} else {
		r0 = ret.Get(0).(pipeline.Result)
	}

	return r0
}

// SetOutputTask provides a mock function with given fields: task
func (_m *Task) SetOutputTask(task pipeline.Task) {
	_m.Called(task)
}

// Type provides a mock function with given fields:
func (_m *Task) Type() pipeline.TaskType {
	ret := _m.Called()

	var r0 pipeline.TaskType
	if rf, ok := ret.Get(0).(func() pipeline.TaskType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pipeline.TaskType)
	}

	return r0
}