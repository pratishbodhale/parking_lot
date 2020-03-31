package vehicle

//go:generate mockgen -destination=mock/vehicle.go -package=mock example.com/parking_lot/vehicle Vehicle
// Interface is named as vehicle than car since in future there might be need to support parking bikes and trucks.
type Vehicle interface {
	RegistrationNumber() string
	Color() string
}

// Unexported since callee doesn't need to have access to struct elements
type car struct {
	registrationNo string
	color          string
}

// Returns new car with given registration number and color
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
