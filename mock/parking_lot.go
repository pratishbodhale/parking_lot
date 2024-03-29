// Code generated by MockGen. DO NOT EDIT.
// Source: example.com/parking_lot (interfaces: ParkingLotManager)

// Package mock is a generated GoMock package.
package mock

import (
	slot "example.com/parking_lot/slot"
	vehicle "example.com/parking_lot/vehicle"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockParkingLotManager is a mock of ParkingLotManager interface
type MockParkingLotManager struct {
	ctrl     *gomock.Controller
	recorder *MockParkingLotManagerMockRecorder
}

// MockParkingLotManagerMockRecorder is the mock recorder for MockParkingLotManager
type MockParkingLotManagerMockRecorder struct {
	mock *MockParkingLotManager
}

// NewMockParkingLotManager creates a new mock instance
func NewMockParkingLotManager(ctrl *gomock.Controller) *MockParkingLotManager {
	mock := &MockParkingLotManager{ctrl: ctrl}
	mock.recorder = &MockParkingLotManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParkingLotManager) EXPECT() *MockParkingLotManagerMockRecorder {
	return m.recorder
}

// FindVehicleSlot mocks base method
func (m *MockParkingLotManager) FindVehicleSlot(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindVehicleSlot", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindVehicleSlot indicates an expected call of FindVehicleSlot
func (mr *MockParkingLotManagerMockRecorder) FindVehicleSlot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVehicleSlot", reflect.TypeOf((*MockParkingLotManager)(nil).FindVehicleSlot), arg0)
}

// LeaveVehicle mocks base method
func (m *MockParkingLotManager) LeaveVehicle(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveVehicle", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaveVehicle indicates an expected call of LeaveVehicle
func (mr *MockParkingLotManagerMockRecorder) LeaveVehicle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveVehicle", reflect.TypeOf((*MockParkingLotManager)(nil).LeaveVehicle), arg0)
}

// ParkVehicle mocks base method
func (m *MockParkingLotManager) ParkVehicle(arg0 vehicle.Vehicle) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParkVehicle", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParkVehicle indicates an expected call of ParkVehicle
func (mr *MockParkingLotManagerMockRecorder) ParkVehicle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParkVehicle", reflect.TypeOf((*MockParkingLotManager)(nil).ParkVehicle), arg0)
}

// SlotsWithColor mocks base method
func (m *MockParkingLotManager) SlotsWithColor(arg0 string) ([]slot.ParkingSlot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlotsWithColor", arg0)
	ret0, _ := ret[0].([]slot.ParkingSlot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SlotsWithColor indicates an expected call of SlotsWithColor
func (mr *MockParkingLotManagerMockRecorder) SlotsWithColor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlotsWithColor", reflect.TypeOf((*MockParkingLotManager)(nil).SlotsWithColor), arg0)
}

// Status mocks base method
func (m *MockParkingLotManager) Status() ([]slot.ParkingSlot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].([]slot.ParkingSlot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockParkingLotManagerMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockParkingLotManager)(nil).Status))
}
