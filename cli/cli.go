package cli

import (
	"errors"
	"strings"
)

type Cli struct {
	cmd command
}

func NewCli() *Cli {
	c := new(Cli)
	c.cmd = command{}
	return c
}

func (c *Cli) Execute(cmd string) (string, error) {
	cmdSlice := strings.Split(cmd, " ")

	execCommand := cmdSlice[0]
	args := cmdSlice[1:]

	switch execCommand {
	case CREATE_PARKING_LOT:
		return c.cmd.createParkingLot(args)
	case PARK:
		return c.cmd.park(args)

	case STATUS:
		return c.cmd.status(args)

	case LEAVE:
		return c.cmd.leave(args)

	case VEHICLES_WITHCOLOR:
		return c.cmd.vehiclesWithColor(args)

	case SLOTS_WITHCOLOR:
		return c.cmd.slotsWithColor(args)

	case SLOT_FOR_REG_NO:
		return c.cmd.slotForRegNo(args)

	default:
		return "", errors.New(NO_CMD)
	}
}
