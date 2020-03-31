package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Unit test check if call to respective function is routed properly
func Test_Execute(t *testing.T) {

	c := NewCli()
	commands := []string{CREATE_PARKING_LOT, PARK, STATUS, LEAVE, VEHICLES_WITHCOLOR, SLOTS_WITHCOLOR, SLOT_FOR_REG_NO}

	for _, cmd := range commands {
		str := cmd + " with invalid args"
		_, err := c.Execute(str)
		assert.NotNil(t, err)

		// Checking if function returns respective error
		assert.EqualValues(t, err.Error(), INVALID_NUMBEROF_ARGUMENTS)
	}

	_, err := c.Execute("invalidCommand")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Error(), NO_CMD)
}
