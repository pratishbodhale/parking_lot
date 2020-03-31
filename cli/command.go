package cli

import (
	"errors"
	"example.com/parking_lot"
	"example.com/parking_lot/vehicle"
	"fmt"
	"strconv"
	"strings"
)

// Unexported since this has use limited to this package
type command struct {
	parkingLot parking_lot.ParkingLotManager
}

// Validate if parking lot manager object is initialized
func (c *command) sanityChecks(argsNeeded, argsGiven int) error {
	if argsGiven != argsNeeded {
		return errors.New(INVALID_NUMBEROF_ARGUMENTS)
	}

	if c.parkingLot == nil {
		return errors.New(PARKING_LOT_NOT_INITIALIZED)
	}
	return nil
}

// If command is create parking lot then initialize the parking lot manager
func (c *command) createParkingLot(cmd []string) (string, error) {
	if len(cmd) != 1 {
		return "", errors.New(INVALID_NUMBEROF_ARGUMENTS)
	}

	noSlots, err := strconv.Atoi(cmd[0])
	if err != nil {
		return "", errors.New(INVALID_ARGUMENT_TYPE)
	}

	c.parkingLot = parking_lot.NewParkingSlotManager(noSlots)

	return fmt.Sprintf("Created a parking lot with %d slots\n", noSlots), nil
}

func (c *command) park(cmd []string) (string, error) {
	if err := c.sanityChecks(2, len(cmd)); err != nil {
		return "", err
	}

	v := vehicle.NewCar(cmd[0], cmd[1])
	slotAssigned, err := c.parkingLot.ParkVehicle(v)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Allocated slot number: %d\n", slotAssigned), nil
}

func (c *command) status(cmd []string) (string, error) {
	if err := c.sanityChecks(0, len(cmd)); err != nil {
		return "", err
	}

	slotsStatus, err := c.parkingLot.Status()
	if err != nil {
		return "", err
	}

	// Standard expoected output doesn't support tab writer

	//const padding = 3
	//w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	//if _, err = fmt.Fprintln(w, "Slot No.\tRegistration No\tColour\t"); err != nil {
	//	return err
	//}

	var out = fmt.Sprintf("%-12s%-19s%s\n", "Slot No.", "Registration No", "Colour")

	for dist, slot := range slotsStatus {
		if free, _ := slot.IsFree(); !free {

			// Ignoring error as We have already checked if slot is not free
			// In current situation there can't be any other error
			v, _ := slot.GetVehicle()
			//if _, err = fmt.Fprintf(w, "%d \t %s \t %s \t\n", dist+1, strings.ToUpper(v.RegistrationNumber()),
			//	strings.Title(v.Color())); err != nil {
			//	return err
			//}
			out += fmt.Sprintf("%-12v%-19v%v\n", dist+1, strings.ToUpper(v.RegistrationNumber()),
				strings.Title(v.Color()))
		}
	}

	//if err = w.Flush(); err != nil {
	//	return err
	//}

	return fmt.Sprint(out), nil
}

func (c *command) leave(cmd []string) (string, error) {
	if err := c.sanityChecks(1, len(cmd)); err != nil {
		return "", err
	}

	noSlot, err := strconv.Atoi(cmd[0])
	if err != nil {
		return "", errors.New(INVALID_ARGUMENT_TYPE)
	}

	err = c.parkingLot.LeaveVehicle(noSlot)
	if err != nil {
		return "", err
	}


	return fmt.Sprintf("Slot number %d is free\n", noSlot), nil
}

func (c *command) vehiclesWithColor(cmd []string) (string, error) {
	if err := c.sanityChecks(1, len(cmd)); err != nil {
		return "", err
	}

	slots, err := c.parkingLot.SlotsWithColor(cmd[0])
	if err != nil {
		return "", err
	}

	var vehicles []string
	for _, slot := range slots {
		v, _ := slot.GetVehicle()
		vehicles = append(vehicles, strings.ToUpper(v.RegistrationNumber()))
	}

	return fmt.Sprintln(strings.Join(vehicles, ", ")), nil
}

func (c *command) slotsWithColor(cmd []string) (string, error) {
	if err := c.sanityChecks(1, len(cmd)); err != nil {
		return "", err
	}

	slots, err := c.parkingLot.SlotsWithColor(cmd[0])
	if err != nil {
		return "", err
	}

	var slotDistances []string
	for _, slot := range slots {
		d, _ := slot.Distance()
		slotDistances = append(slotDistances, strconv.Itoa(d))
	}

	return fmt.Sprintln(strings.Join(slotDistances, ", ")), nil
}

func (c *command) slotForRegNo(cmd []string) (string, error) {
	if err := c.sanityChecks(1, len(cmd)); err != nil {
		return "", err
	}

	s, err := c.parkingLot.FindVehicleSlot(cmd[0])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d\n", s), nil
}
