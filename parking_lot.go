package parking_lot

type Vehicle interface {
	RegistrationNumber() string
	Color() string
}

type ParkingSlot interface {
	IsFree() (bool, error)
	Free() error
	GetVehicle() (Vehicle, error)
	ParkVehicle(Vehicle) error
	Distance() (int, error)
}

type ParkingLotManager interface {
	Init(slots int) error
	ParkVehicle(v Vehicle) (int, error)
	FindVehicleSlot(registrationNumber string) (int, error)
	LeaveVehicle(s int) error
	Status() ([]ParkingSlot, error)
	VehiclesWithColor(color string) ([]Vehicle, error)
}