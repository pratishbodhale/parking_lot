package slot

import (
	"errors"
	"example.com/parking_lot/vehicle"
)

type ParkingSlot interface {
	IsFree() (bool, error)
	Free() error
	GetVehicle() (vehicle.Vehicle, error)
	ParkVehicle(v vehicle.Vehicle) error
	Distance() (int, error)
}

const(
	NO_VEHICLE = "No vehicle is parked in slot "
	NOT_FREE = "slot is not free "
) 

type parkingSlotImpl struct {
	v        vehicle.Vehicle
	distance int
}

func NewParkingSlot(dist int) ParkingSlot {
	p := parkingSlotImpl{
		v:        nil,
		distance: dist,
	}

	return &p
}

func (p *parkingSlotImpl) IsFree() (bool, error) {
	if p.v == nil {
		return true, nil
	}
	return false, nil
}

func (p *parkingSlotImpl) Free() error {
	p.v = nil
	return nil
}

func (p *parkingSlotImpl) GetVehicle() (vehicle.Vehicle, error) {
	if p.v == nil {
		return nil, errors.New(NO_VEHICLE)
	}

	return p.v, nil
}

func (p *parkingSlotImpl) ParkVehicle(v vehicle.Vehicle) error {
	if p.v != nil {
		return errors.New(NOT_FREE)
	}

	p.v = v
	return nil
}

func (p *parkingSlotImpl) Distance() (int, error) {
	return p.distance, nil
}
