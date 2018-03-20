// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// ExtractedCredential is an autogenerated mock type for the ExtractedCredential type
type ExtractedCredential struct {
	mock.Mock
}

// CreateExtractedCredential provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *ExtractedCredential) CreateExtractedCredential(_a0 string, _a1 string, _a2 map[string]interface{}, _a3 map[string]string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}, map[string]string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteExtractedCredential provides a mock function with given fields: _a0, _a1
func (_m *ExtractedCredential) DeleteExtractedCredential(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetExtractedCredential provides a mock function with given fields: _a0, _a1
func (_m *ExtractedCredential) GetExtractedCredential(_a0 string, _a1 string) (map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string, string) map[string]interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExtractedCredential provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *ExtractedCredential) UpdateExtractedCredential(_a0 string, _a1 string, _a2 map[string]interface{}, _a3 map[string]string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}, map[string]string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
