package vehicle

//go:generate mockgen -destination=mock/vehicle.go -package=mock example.com/parking_lot/vehicle Vehicle
type Vehicle interface {
	RegistrationNumber() string
	Color() string
}

type car struct {
	registrationNo string
	color          string
}

func NewCar(registrationNo, color string) Vehicle {
	c := car{
		registrationNo: registrationNo,
		color:          color,
	}
	return &c
}

func (c *car) RegistrationNumber() string {
	return c.registrationNo
}

func (c *car) Color() string {
	return c.color
}
