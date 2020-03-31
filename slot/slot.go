package slot

import (
	"errors"
	"example.com/parking_lot/vehicle"
)

//go:generate mockgen -destination=mock/slot.go -package=mock example.com/parking_lot/slot ParkingSlot
// Manager a single parking slot
type ParkingSlot interface {
	IsFree() (bool, error)
	Free() error
	GetVehicle() (vehicle.Vehicle, error)
	ParkVehicle(v vehicle.Vehicle) error
	Distance() (int, error)
}

const (
	NO_VEHICLE = "No vehicle is parked in slot "
	NOT_FREE   = "slot is not free "
)

// Unexported since callee doesn't need to have access to struct elements
type parkingSlotImpl struct {
	v        vehicle.Vehicle
	distance int
}

// Creates a new parking slot having specified distance from the entry gate and return it
func NewParkingSlot(dist int) ParkingSlot {
	p := parkingSlotImpl{
		v:        nil,
		distance: dist,
	}

	return &p
}

// Return true if the parking lot is free
func (p *parkingSlotImpl) IsFree() (bool, error) {
	if p.v == nil {
		return true, nil
	}
	return false, nil
}

// Frees the slot
func (p *parkingSlotImpl) Free() error {
	p.v = nil
	return nil
}

// Returns vehicle in the slots, Errors out with not found if slot is free
func (p *parkingSlotImpl) GetVehicle() (vehicle.Vehicle, error) {
	if p.v == nil {
		return nil, errors.New(NO_VEHICLE)
	}

	return p.v, nil
}

// Parks vehicle to the slot
func (p *parkingSlotImpl) ParkVehicle(v vehicle.Vehicle) error {
	if p.v != nil {
		return errors.New(NOT_FREE)
	}

	p.v = v
	return nil
}

// Gets distance from the entry gate
func (p *parkingSlotImpl) Distance() (int, error) {
	return p.distance, nil
}
