package parking_lot

import (
	"errors"
	"example.com/parking_lot/slot"
	"example.com/parking_lot/slot/mock"
	"example.com/parking_lot/vehicle"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewParkingSlotManager(t *testing.T) {

	numSlots := 2
	pm := NewParkingSlotManager(numSlots)
	slots, err := pm.Status()

	assert.Nil(t, err)
	assert.EqualValues(t, len(slots), numSlots)
}

func TestParkingLotMgrImpl_FindVehicleSlot(t *testing.T) {

	v := vehicle.NewCar("rno", "blue")

	ctrl := gomock.NewController(t)
	slotMock := mock.NewMockParkingSlot(ctrl)
	someUnexpError := errors.New("Some unknown error ")
	expDistance := 1

	count := -1
	slotMock.EXPECT().GetVehicle().DoAndReturn(func() (vehicle.Vehicle, error) {
		count++
		if count == 0 {
			return v, errors.New(slot.NO_VEHICLE)
		} else if count == 1 {
			return v, someUnexpError
		}
		return v, nil
	}).Times(3)

	slotMock.EXPECT().Distance().DoAndReturn(func() (int, error) {
		return expDistance, nil
	}).Times(1)

	slots := []slot.ParkingSlot{slotMock}
	pm := parkingLotMgrImpl{ slots}

	// CASE: No vehicles in slot
	_, err := pm.FindVehicleSlot("rno2")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), NOT_FOUND)

	// CASE: Some unexpected error from slots while getting vehicle
	_, err = pm.FindVehicleSlot("rno2")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), someUnexpError.Error())

	// CASE: Vehicle found
	actualV, err := pm.FindVehicleSlot("rno")
	assert.Nil(t, err)
	assert.EqualValues(t, expDistance, actualV)
}

func TestParkingLotMgrImpl_LeaveVehicle(t *testing.T) {
	ctrl := gomock.NewController(t)
	slotMock := mock.NewMockParkingSlot(ctrl)

	slotMock.EXPECT().Free().DoAndReturn(func() error{
		return nil
	}).Times(1)

	slots := []slot.ParkingSlot{slotMock}
	pm := parkingLotMgrImpl{ slots}

	// CASE: Invalid input
	err := pm.LeaveVehicle(0)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), SLOT_OUT_OF_BOUND)

	// CASE: Success
	err = pm.LeaveVehicle(1)
	assert.Nil(t, err)
}

func TestParkingLotMgrImpl_ParkVehicle(t *testing.T) {

	v := vehicle.NewCar("rno", "blue")

	ctrl := gomock.NewController(t)
	slotMock := mock.NewMockParkingSlot(ctrl)

	countIsFree := -1
	slotMock.EXPECT().IsFree().DoAndReturn(func() (bool, error) {
		countIsFree++
		if countIsFree == 0{
			return false, nil
		}
		return true, nil
	}).Times(3)

	countPark := -1
	errParking := errors.New("Unable to park vehicle ")
	slotMock.EXPECT().ParkVehicle(gomock.Any()).DoAndReturn(func(vehicle.Vehicle) error{
		countPark++
		if countPark == 0{
			return errParking
		}
		return nil
	}).Times(2)

	expSlot:= 1
	slotMock.EXPECT().Distance().DoAndReturn(func() (int, error) {
		return expSlot, nil
	}).Times(1)

	slots := []slot.ParkingSlot{slotMock}
	pm := parkingLotMgrImpl{ slots}

	// CASE: No slots free
	_, err := pm.ParkVehicle(v)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), FULL_PARKING)

	// CASE: Unable to park
	_, err = pm.ParkVehicle(v)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), errParking.Error())

	// CASE: Success
	actualSlot , err := pm.ParkVehicle(v)
	assert.Nil(t, err)
	assert.EqualValues(t, expSlot, actualSlot)
}

func TestParkingLotMgrImpl_SlotsWithColor(t *testing.T) {

	v := vehicle.NewCar("rno", "blue")

	ctrl := gomock.NewController(t)
	slotMock := mock.NewMockParkingSlot(ctrl)
	someUnexpError := errors.New("Some unknown error ")
	expDistance := 1

	count := -1
	slotMock.EXPECT().GetVehicle().DoAndReturn(func() (vehicle.Vehicle, error) {
		count++
		if count == 0 {
			return v, errors.New(slot.NO_VEHICLE)
		} else if count == 1 {
			return v, someUnexpError
		}
		return v, nil
	}).Times(3)

	slotMock.EXPECT().Distance().DoAndReturn(func() (int, error) {
		return expDistance, nil
	}).Times(1)

	slots := []slot.ParkingSlot{slotMock}
	pm := parkingLotMgrImpl{ slots}

	// CASE: No vehicles in slot
	_, err := pm.SlotsWithColor("blue")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), NOT_FOUND)

	// CASE: Some unexpected error from slots while getting vehicle
	_, err = pm.SlotsWithColor("blue")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), someUnexpError.Error())

	// CASE: Vehicle found
	actualV, err := pm.SlotsWithColor("blue")
	assert.Nil(t, err)
	assert.EqualValues(t, slots, actualV)
}