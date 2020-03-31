package parking_lot

import (
	"errors"
	"example.com/parking_lot/slot"
	"example.com/parking_lot/vehicle"
)

//go:generate mockgen -destination=mock/parking_lot.go -package=mock example.com/parking_lot ParkingLotManager

// Manages all operations of parking lot
type ParkingLotManager interface {
	ParkVehicle(v vehicle.Vehicle) (int, error)
	FindVehicleSlot(registrationNumber string) (int, error)
	LeaveVehicle(s int) error
	Status() ([]slot.ParkingSlot, error)
	SlotsWithColor(color string) ([]slot.ParkingSlot, error)
}

const (
	FULL_PARKING      = "Sorry, parking lot is full"
	NOT_FOUND         = "Not found"
	SLOT_OUT_OF_BOUND = "Invalid error, Slot out bound"
)

// Saves pointer to all the slots in parking lot
// Unexported since callee doesn't need to have access to struct elements
type parkingLotMgrImpl struct {
	slots []slot.ParkingSlot
}

// Returns a new parking lot manager with given number of slots
func NewParkingSlotManager(slots int) ParkingLotManager {
	p := parkingLotMgrImpl{}
	p.init(slots)
	return &p
}

func (p *parkingLotMgrImpl) init(slots int) error {

	for i := 0; i < slots; i++ {
		s := slot.NewParkingSlot(i + 1)
		p.slots = append(p.slots, s)
	}

	return nil
}

// Parks vehicle to the nearest possible slot
// Return full parking error if not slots found
func (p *parkingLotMgrImpl) ParkVehicle(v vehicle.Vehicle) (int, error) {
	for _, s := range p.slots {
		if isFree, err := s.IsFree(); err == nil && isFree {
			if err = s.ParkVehicle(v); err != nil {
				return -1, err
			}
			return s.Distance()
		}
	}
	return -1, errors.New(FULL_PARKING)
}

// Returns slot where vehicle with a registration number is present
func (p *parkingLotMgrImpl) FindVehicleSlot(registrationNumber string) (int, error) {
	for _, s := range p.slots {
		v, err := s.GetVehicle()

		// If err is that slot doesn't have a vehicle
		if err != nil && err.Error() == slot.NO_VEHICLE {
			continue

		// Some other error like database load
		} else if err != nil {
			return -1, err
		}

		// Vehicle found
		if v.RegistrationNumber() == registrationNumber {
			return s.Distance()
		}
	}
	return -1, errors.New(NOT_FOUND)
}

// Unparks the vehicle
func (p *parkingLotMgrImpl) LeaveVehicle(s int) error {
	if len(p.slots) < s || s < 1 {
		return errors.New(SLOT_OUT_OF_BOUND)
	}
	return p.slots[s-1].Free()
}

// Returns status of all slots
// Vehicle details can be accesed using slot methods
func (p *parkingLotMgrImpl) Status() ([]slot.ParkingSlot, error) {
	return p.slots, nil
}

// Returns all slots that has vehicle of a given color
// Vehicle details can be accesed using slot methods
func (p *parkingLotMgrImpl) SlotsWithColor(color string) ([]slot.ParkingSlot, error) {
	var slots []slot.ParkingSlot

	for _, s := range p.slots {
		v, err := s.GetVehicle()

		// If err is that slot doesn't have a vehicle
		if err != nil && err.Error() == slot.NO_VEHICLE {
			continue

			// Some other error like database load
		} else if err != nil {
			return nil, err
		}

		if v.Color() == color {
			slots = append(slots, s)
		}
	}
	if len(slots) == 0 {
		return nil, errors.New(NOT_FOUND)
	}
	return slots, nil
}
