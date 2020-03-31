package slot

import (
	"example.com/parking_lot/vehicle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParkingSlotImpl_Distance(t *testing.T) {
	dist := 1
	slot := NewParkingSlot(dist)
	expDist, err := slot.Distance()
	assert.Nil(t, err)
	assert.EqualValues(t, expDist, dist)
}

func TestParkingSlotImpl_ParkVehicle(t *testing.T) {
	dist := 1
	slot := NewParkingSlot(dist)

	v := vehicle.NewCar("rno", "blue")

	var err error

	// Success parking vehicle
	err = slot.ParkVehicle(v)
	assert.Nil(t, err)

	// Failure parking vehicle
	err = slot.ParkVehicle(v)
	assert.NotNil(t, err)
}

func TestParkingSlotImpl_IsFree(t *testing.T) {
	dist := 1
	slot := NewParkingSlot(dist)

	v := vehicle.NewCar("rno", "blue")

	var err error

	// Success parking vehicle
	err = slot.ParkVehicle(v)
	assert.Nil(t, err)

	// Checking if free should return false
	isFree, err := slot.IsFree()
	assert.Nil(t, err)
	assert.EqualValues(t, isFree, false)

	// Free the slot
	err = slot.Free()
	assert.Nil(t, err)

	// Checking if free should return true
	isFree, err = slot.IsFree()
	assert.Nil(t, err)
	assert.EqualValues(t, isFree, true)
}

func TestParkingSlotImpl_GetVehicle(t *testing.T) {
	dist := 1
	slot := NewParkingSlot(dist)

	var err error

	// Get vehicle it is not parked
	_, err = slot.GetVehicle()
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), NO_VEHICLE)

	expV := vehicle.NewCar("rno", "blue")

	// Success parking vehicle
	err = slot.ParkVehicle(expV)
	assert.Nil(t, err)

	// Get vehicle when it is parked
	actualV, err := slot.GetVehicle()
	assert.Nil(t, err)
	assert.EqualValues(t, actualV, expV)

}
