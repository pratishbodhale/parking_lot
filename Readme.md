# Parking Lot Problem

When a car enters a parking lot a ticket issued to the driver. 
A nearest available parking slot is alloted to the vehicle before actually handing over a ticket.
Vehicle can be identified using registration number (number plate) and the colour.
At the exit the customer returns the ticket which then marks the slot they were using as being available.

## Setting up application

Execute ```setup``` in ```bin``` directory. This command runs all the unit & functional tests and creates binary ```parking_lot``` in bin directory.

```sh
parking_lot $ ./bin/setup
```

Application can be run by executing bash command ```parking_lot``` in bin directory with 2 options:

* Inputs taken from file specified
```sh
parking_lot $ ./bin/parking_lot [input_filepath]
```
* Inputs are taken from command prompt.
```sh
parking_lot $ ./bin/parking_lot
```

## Command list ##

* ```create_parking_lot [capacity]```
Initialization of parking lot with parameters of slot capacity. 

* ```park [car_registration_number] [car_color]```
Park creates a car object with given registration number and car color. Parks in one of the slot available.

* ```leave [slot_number]```
Vacates the slot taken by car at slot number

* ```status```
Prints the current status of all slots

* ```registration_numbers_for_cars_with_colour [car_color]```
Prints all the cars registrations numbers with given color

* ```slot_numbers_for_cars_with_colour [car_color]```
Prints all the slots taken by vehicle with given color

* ```slot_number_for_registration_number [car_registration_number]```
Prints slot for a car with given registration number


## Application Components

#### Vehicle
Interfaces name is kept as vehicle than car since, In future there can be trucks, bikes for which we support in the future.

```go
type Vehicle interface {
	RegistrationNumber() string
	Color() string
}
```

There is an object with name car. It can be constructed by using ```NewCar()``` function.
Car object has ```registrationNo``` and ```color``` properties


#### Parking Slot
Interface slot is created since in future there can be multiple slots with different sizes for different kinds of vehicles.
Most functions return values along with error since in future this application might have connection to database and other depencies where failure might occur.
Those all should just implement this interface. Minimal changes will be needed to adapt to that design.

```go
type ParkingSlot interface {
	IsFree() (bool, error)
	Free() error
	GetVehicle() (vehicle.Vehicle, error)
	ParkVehicle(v vehicle.Vehicle) error
	Distance() (int, error)
}
```

Object implementation of this stores the vehicle in that slot (```v```) and all stores how far that slot is for the gate (```distance```).

#### Parking lot manager
It is a main component that has access to all the slots

```go
type ParkingLotManager interface {
	ParkVehicle(v vehicle.Vehicle) (int, error)
	FindVehicleSlot(registrationNumber string) (int, error)
	LeaveVehicle(s int) error
	Status() ([]slot.ParkingSlot, error)
	SlotsWithColor(color string) ([]slot.ParkingSlot, error)
}
```

Object implementation saves pointer to all the slots in the parking lot (```slots```). 

## Unit testing
Since the application is designed using interfaces it is easy to mock and test each package independent of other package's behaviour.
Each package contains respective tests for all the functions. Code coverage is 100% except the main function. 
Used ```gomock``` package to mock the interfaces and ```assert``` package to equate the expected results. Readability can improved if tool like goconvey is used.