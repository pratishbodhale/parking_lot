package vehicle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCar(t *testing.T) {
	registrationNo := "rno"
	color := "blue"
	car := NewCar(registrationNo, color)

	assert.Equal(t, car.RegistrationNumber(), registrationNo)

	assert.Equal(t, car.Color(), color)
}
