package cli

import (
	"errors"
	"example.com/parking_lot/mock"
	"example.com/parking_lot/slot"
	"example.com/parking_lot/vehicle"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cmd []string
var err error

func TestCommand_CreateParkingLot(t *testing.T) {
	var c = command{}

	cmd = []string{"2", "1"}
	_, err = c.createParkingLot(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)

	cmd = []string{"ab"}
	_, err = c.createParkingLot(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_ARGUMENT_TYPE)

	cmd = []string{"2"}
	_, err = c.createParkingLot(cmd)
	assert.Nil(t, err)
}

func TestCommand_Park(t *testing.T) {
	var c = command{}

	ctrl := gomock.NewController(t)
	pm := mock.NewMockParkingLotManager(ctrl)

	errUnablePark := errors.New("Unable to park vehicle ")
	count := -1
	pm.EXPECT().ParkVehicle(gomock.Any()).DoAndReturn(func(vehicle vehicle.Vehicle) (int, error) {
		count++
		if count == 0 {
			return -1, errUnablePark
		}
		return 1, nil

	}).Times(2)

	// Call without initializing parking lot manager
	cmd = []string{"regNo", "blue"}
	_, err = c.park(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), PARKING_LOT_NOT_INITIALIZED)

	c.parkingLot = pm

	cmd = []string{"reg"}
	_, err = c.park(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)

	cmd = []string{"regNo", "blue"}
	_, err = c.park(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), errUnablePark.Error())

	cmd = []string{"regNo", "blue"}
	_, err = c.park(cmd)
	assert.Nil(t, err)
}

func TestCommand_Status(t *testing.T) {
	var c = command{}

	ctrl := gomock.NewController(t)
	pm := mock.NewMockParkingLotManager(ctrl)

	errUnableStatus := errors.New("Unable to get status ")

	v := vehicle.NewCar("regNo", "blue")
	sl := slot.NewParkingSlot(1)
	err := sl.ParkVehicle(v)
	assert.Nil(t, err)

	expResp := []slot.ParkingSlot{sl}
	count := -1
	pm.EXPECT().Status().DoAndReturn(func() ([]slot.ParkingSlot, error) {
		count++
		if count == 0 {
			return nil, errUnableStatus
		}
		return expResp, nil

	}).Times(2)

	c.parkingLot = pm

	cmd = []string{"reg"}
	_, err = c.status(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)

	cmd = []string{}
	_, err = c.status(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), errUnableStatus.Error())

	cmd = []string{}
	_, err = c.status(cmd)
	assert.Nil(t, err)
}

func TestCommand_Leave(t *testing.T) {
	var c = command{}

	ctrl := gomock.NewController(t)
	pm := mock.NewMockParkingLotManager(ctrl)

	errUnableUnPark := errors.New("Unable to park vehicle ")
	count := -1
	pm.EXPECT().LeaveVehicle(gomock.Any()).DoAndReturn(func(slot int) error {
		count++
		if count == 0 {
			return errUnableUnPark
		}
		return nil

	}).Times(2)

	c.parkingLot = pm

	cmd = []string{"reg", "1"}
	_, err = c.leave(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)

	cmd = []string{"blue"}
	_, err = c.leave(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_ARGUMENT_TYPE)

	cmd = []string{"1"}
	_, err = c.leave(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), errUnableUnPark.Error())

	cmd = []string{"1"}
	_, err = c.leave(cmd)
	assert.Nil(t, err)
}

func TestCommand_VehiclesWithColor(t *testing.T) {
	var c = command{}

	ctrl := gomock.NewController(t)
	pm := mock.NewMockParkingLotManager(ctrl)

	errUnable := errors.New("Unable to get slots ")

	v := vehicle.NewCar("regNo", "blue")
	sl := slot.NewParkingSlot(1)
	err := sl.ParkVehicle(v)
	assert.Nil(t, err)

	expResp := []slot.ParkingSlot{sl}
	count := -1
	pm.EXPECT().SlotsWithColor(gomock.Any()).DoAndReturn(func(color string) ([]slot.ParkingSlot, error) {
		count++
		if count == 0 {
			return nil, errUnable
		}
		return expResp, nil

	}).Times(2)

	c.parkingLot = pm

	cmd = []string{"reg", "reg"}
	_, err = c.vehiclesWithColor(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)

	cmd = []string{"blue"}
	_, err = c.vehiclesWithColor(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), errUnable.Error())

	cmd = []string{"blue"}
	_, err = c.vehiclesWithColor(cmd)
	assert.Nil(t, err)
}

func TestCommand_SlotsWithColor(t *testing.T) {
	var c = command{}

	ctrl := gomock.NewController(t)
	pm := mock.NewMockParkingLotManager(ctrl)

	errUnable := errors.New("Unable to get slots ")

	v := vehicle.NewCar("regNo", "blue")
	sl := slot.NewParkingSlot(1)
	err := sl.ParkVehicle(v)
	assert.Nil(t, err)

	expResp := []slot.ParkingSlot{sl}
	count := -1
	pm.EXPECT().SlotsWithColor(gomock.Any()).DoAndReturn(func(color string) ([]slot.ParkingSlot, error) {
		count++
		if count == 0 {
			return nil, errUnable
		}
		return expResp, nil

	}).Times(2)

	c.parkingLot = pm

	cmd = []string{"reg", "reg"}
	_, err = c.slotsWithColor(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)

	cmd = []string{"blue"}
	_, err = c.slotsWithColor(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), errUnable.Error())

	cmd = []string{"blue"}
	_, err = c.slotsWithColor(cmd)
	assert.Nil(t, err)
}

func TestCommand_SlotForRegNo(t *testing.T) {
	var c = command{}

	ctrl := gomock.NewController(t)
	pm := mock.NewMockParkingLotManager(ctrl)

	errResp := errors.New("Unable to get status ")

	v := vehicle.NewCar("regNo", "blue")
	sl := slot.NewParkingSlot(1)
	err := sl.ParkVehicle(v)
	assert.Nil(t, err)

	count := -1
	pm.EXPECT().FindVehicleSlot(gomock.Any()).DoAndReturn(func(regNo string) (int, error) {

		count++
		if count == 0 {
			return -1, errResp
		}
		return 1, nil

	}).Times(2)

	c.parkingLot = pm

	cmd = []string{"reg", "reg"}
	_, err = c.slotForRegNo(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)

	cmd = []string{"reg"}
	_, err = c.slotForRegNo(cmd)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), errResp.Error())

	cmd = []string{"reg"}
	_, err = c.slotForRegNo(cmd)
	assert.Nil(t, err)
}
