package parking_lot

import (
	"errors"
	"example.com/parking_lot/slot"
	"example.com/parking_lot/vehicle"
)

type ParkingLotManager interface {
	ParkVehicle(v vehicle.Vehicle) (int, error)
	FindVehicleSlot(registrationNumber string) (int, error)
	LeaveVehicle(s int) error
	Status() ([]slot.ParkingSlot, error)
	SlotsWithColor(color string) ([]slot.ParkingSlot, error)
}

const (
	UNABLE_TO_PARK = "Unable to park the vehicle "
	FULL_PARKING = "Parking is full "
	NO_VEHICLE_FOUND = "No vehicle with registration number found "
)

type parkingLotMgrImpl struct {
	slots []slot.ParkingSlot
}

func NewParkingSlotManager(slots int) ParkingLotManager {
	p := parkingLotMgrImpl{}
	p.init(slots)
	return &p
}

func (p *parkingLotMgrImpl) init(slots int) error {

	for i := 0; i < slots; i++ {
		s := slot.NewParkingSlot(i+1)
		p.slots = append(p.slots, s)
	}

	return nil
}

func (p *parkingLotMgrImpl) ParkVehicle(v vehicle.Vehicle) (int, error){
	for _, s := range p.slots{
		if isFree, _ := s.IsFree(); isFree {
			if err := s.ParkVehicle(v); err != nil {
				return -1, errors.New(UNABLE_TO_PARK)
			}
			return s.Distance()
		}
	}
	return -1, errors.New(FULL_PARKING)
}

func (p *parkingLotMgrImpl) FindVehicleSlot(registrationNumber string) (int, error){
	for _, s := range p.slots{
		if isFree, _ := s.IsFree(); !isFree {

			v, _ := s.GetVehicle()
			if v.RegistrationNumber() == registrationNumber{
				return s.Distance()
			}
		}
	}
	return -1, errors.New(NO_VEHICLE_FOUND)
}

func (p *parkingLotMgrImpl) LeaveVehicle(s int) error{
	return p.slots[s-1].Free()
}

func (p *parkingLotMgrImpl) Status() ([]slot.ParkingSlot, error){
	return p.slots, nil
}

func (p *parkingLotMgrImpl) SlotsWithColor(color string) ([]slot.ParkingSlot, error){
	var slots []slot.ParkingSlot
	for _, s := range p.slots{
		if isFree, _ := s.IsFree(); !isFree {

			v, _ := s.GetVehicle()
			if v.Color() == color{
				slots = append(slots, s)
			}
		}
	}
	return slots, nil
}

