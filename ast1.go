package main

import (
	"fmt"
	"errors"
)

type Vehicle struct{
	Brand string
	Model string
	Year int
	Color string
}

type VehicleError struct{
	VehicleType string
	Err error
}


func (ve VehicleError) Error() string{
	return fmt.Sprintf("%s vehicle error: %v", ve.VehicleType, ve.Err)
}

type VehicleInterface interface{
	Start()
	Stop()
	Steer()
}

type Car struct{
	Vehicle
	NumDoors int
	EngineType string
}

type Boat struct{
	Vehicle
	Length int
	PropulsionType string
}

type MotorCycle struct{
	Vehicle
	NumWheels int
	HasSideCar bool
}

func (c Car) Start(){
	fmt.Printf("Starting the %s %s Car with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

func (c Car) Stop(){
	fmt.Printf("Stopping the %s %s Car with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

func (c Car) Steer(){
	fmt.Printf("Steering the %s %s Car with %d doors\n", c.Color, c.Brand, c.NumDoors)
}


func (b Boat) Start(){
	fmt.Printf("Starting the %s boat with %d length\n", b.Brand, b.Length)
}

func (b Boat) Stop(){
	fmt.Printf("Stopping the %s boat with %d length\n", b.Brand, b.Length)
}

func (b Boat) Steer(){
	fmt.Printf("Steering the %s boat with %d length\n", b.Brand, b.Length)
}

func (m MotorCycle) Start(){
	fmt.Printf("Starting the %s motorcycle with %d wheels\n", m.Brand, m.NumWheels)
	if m.NumWheels < 0 {
		panic("negative number of wheels")
	}
}

func (m MotorCycle) Stop(){
	fmt.Printf("Stopping the %s motorcycle with %d wheels\n", m.Brand, m.NumWheels)
}

func (m MotorCycle) Steer(){
	fmt.Printf("Steering the %s motorcycle with %d wheels\n", m.Brand, m.NumWheels)
}



func AddVehicle(vehicleType string, brand string, model string, year int, color string) (VehicleInterface, error){
	switch vehicleType {
	case "car":
		return Car{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, 4, "gasoline"}, nil
	case "boat":
		return Boat{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, 20, "motor"}, nil
	// case "motorcycle":
	// 	return MotorCycle{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, 2, false}, nil
	case "motorcycle":
		return MotorCycle{Vehicle{Brand: brand, Model: model, Year: year, Color: color}, -1, false}, nil // simulate an error
	default:
		return nil, VehicleError{VehicleType: "Unknown", Err: errors.New("Invalid Vehicle Type ")}
	}
}


func VehicleAbilities(v VehicleInterface){
	v.Start()
	v.Steer()
	v.Stop()
}

func doRecover(){
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main(){
	
	defer doRecover()

	car, err := AddVehicle("car", "Toyota", "Camry", 2020, "blue")
	if err != nil {
		fmt.Println("Error adding car:", err)
		return
	}

	boat, err := AddVehicle("boat", "bayliner", "element", 2018, "white")
	if err != nil {
		fmt.Println("error adding boat:", err)
		return
	}

	motorcycle, err := AddVehicle("motorcycle", "harley", "street glide", 2022, "black")
	if err != nil {
		fmt.Println("error adding motorcycle:", err)
		return
	}

	VehicleAbilities(car)
	VehicleAbilities(boat)
	VehicleAbilities(motorcycle)

}