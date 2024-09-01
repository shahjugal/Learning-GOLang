package main

import "fmt"

// unlike other language here we dont use nouns as interfaces but verbs as interfaces.

type Fuelable interface {
	FuelUp()
}

type Rideable interface {
	Ride()
}

type car struct {
}

type bike struct {
}

func (c *car) FuelUp() {
	fmt.Println("Fueled the car")
}

func (c *car) Ride() {
	fmt.Println("Riding the car")
}

func (b *bike) FuelUp() {
	fmt.Println("Fueled the bike")
}

func (b *bike) Ride() {
	fmt.Println("Riding the bike")
}

func fuelVehicle(f Fuelable) {
	f.FuelUp()
}

func rideVehicle(r Rideable) {
	r.Ride()
}

func main() {

	fuelVehicle(&bike{})
	rideVehicle(&bike{})

	fuelVehicle(&car{})
	rideVehicle(&car{})

}
