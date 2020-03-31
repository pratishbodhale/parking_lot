package cli

import (
	"errors"
	"example.com/parking_lot"
	"example.com/parking_lot/vehicle"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

// Commands
const (
	CREATE_PARKING_LOT = "create_parking_lot"
	PARK = "park"
	LEAVE = "leave"
	STATUS = "status"
	VEHICLES_WITHCOLOR = "registration_numbers_for_cars_with_colour"
	SLOTS_WITHCOLOR = "slot_numbers_for_cars_with_colour"
	SLOT_FOR_REG_NO = "slot_number_for_registration_number"
)

const (
	NO_CMD = "No command exists "
	INVALID_NUMBEROF_ARGUMENTS = "Invalid number of arguments "
	INVALID_ARGUMENT_TYPE = "Invalid arguments "
	PARKING_LOT_NOT_INITIALIZED = "Parking lot not initialized "
)

func (c *Cli) isParkingLotManagerInit() bool {
	return c.parkingLot != nil
}

func (c *Cli) execute(cmd string) error {
	cmdSlice := strings.Split(cmd, " ")
	cmdLen := len(cmdSlice)

	if cmdLen < 1 {
		return errors.New(NO_CMD)
	}

	switch cmdSlice[0] {
	case CREATE_PARKING_LOT:
		if cmdLen != 2{
			return errors.New(INVALID_NUMBEROF_ARGUMENTS)
		}

		noSlots, err := strconv.Atoi(cmdSlice[1])
		if err != nil {
			return errors.New(INVALID_ARGUMENT_TYPE)
		}

		c.parkingLot = parking_lot.NewParkingSlotManager(noSlots)
		fmt.Printf("Created a parking lot with %d slots\n", noSlots)

	case PARK:
		if cmdLen != 3 {
			return  errors.New(INVALID_NUMBEROF_ARGUMENTS)
		} else if c.isParkingLotManagerInit() == false{
			return  errors.New(PARKING_LOT_NOT_INITIALIZED)
		}

		v := vehicle.NewCar(cmdSlice[1], cmdSlice[2])
		slotAssigned, err := c.parkingLot.ParkVehicle(v)
		if err!=nil{
			return err
		}
		fmt.Printf("Allocated slot number: %d\n", slotAssigned)

	case STATUS:
		if cmdLen != 1{
			return  errors.New(INVALID_NUMBEROF_ARGUMENTS)
		}else if c.isParkingLotManagerInit() == false{
			return  errors.New(PARKING_LOT_NOT_INITIALIZED)
		}

		slotsStatus, err := c.parkingLot.Status()
		if err != nil{
			return err
		}

		const padding = 3
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
		fmt.Fprintln(w, "Slot No.\tRegistration No\tColour\t")
		for dist, slot := range slotsStatus{
			if free, _ := slot.IsFree(); !free {

				// Ignoring error as We have already checked if slot is not free
				// In current situation there can't be any other error
				v, _  := slot.GetVehicle()
				if _, err = fmt.Fprintf(w, "%d \t %s \t %s \t\n", dist+1 , strings.ToUpper(v.RegistrationNumber()),
					strings.Title(v.Color())); err != nil{
						return err
				}

			}
		}
		w.Flush()

	case LEAVE:
		if cmdLen != 2{
			return  errors.New(INVALID_NUMBEROF_ARGUMENTS)
		}else if c.isParkingLotManagerInit() == false{
			return  errors.New(PARKING_LOT_NOT_INITIALIZED)
		}

		noSlot, err := strconv.Atoi(cmdSlice[1])
		if err != nil {
			return errors.New(INVALID_ARGUMENT_TYPE)
		}

		err = c.parkingLot.LeaveVehicle(noSlot)
		if err != nil{
			return err
		}

		fmt.Printf("Slot number %d is free\n", noSlot)

	case VEHICLES_WITHCOLOR:
		if cmdLen != 2{
			return  errors.New(INVALID_NUMBEROF_ARGUMENTS)
		}else if c.isParkingLotManagerInit() == false{
			return  errors.New(PARKING_LOT_NOT_INITIALIZED)
		}

		slots, err := c.parkingLot.SlotsWithColor(cmdSlice[1])
		if err != nil {
			return err
		}

		var vehicles []string
		for _, slot := range slots{
			v, _ := slot.GetVehicle()
			vehicles = append(vehicles, strings.ToUpper(v.RegistrationNumber()))
		}
		fmt.Println(strings.Join(vehicles, ", "))

	case SLOTS_WITHCOLOR:
		if cmdLen != 2{
			return  errors.New(INVALID_NUMBEROF_ARGUMENTS)
		}else if c.isParkingLotManagerInit() == false{
			return  errors.New(PARKING_LOT_NOT_INITIALIZED)
		}

		slots, err := c.parkingLot.SlotsWithColor(cmdSlice[1])
		if err != nil {
			return err
		}

		var slotDistances []string
		for _, slot := range slots{
			d, _ := slot.Distance()
			slotDistances = append(slotDistances, strconv.Itoa(d))
		}
		fmt.Println(strings.Join(slotDistances, ", "))

	case SLOT_FOR_REG_NO:
		if cmdLen != 2{
			return  errors.New(INVALID_NUMBEROF_ARGUMENTS)
		}else if c.isParkingLotManagerInit() == false{
			return  errors.New(PARKING_LOT_NOT_INITIALIZED)
		}
		s, err := c.parkingLot.FindVehicleSlot(cmdSlice[1])
		if err != nil{
			return err
		}
		fmt.Printf("%d \n", s)

	default:
		fmt.Println("No command found")
	}

	return nil
}
