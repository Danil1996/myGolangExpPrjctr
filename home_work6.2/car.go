package home_work6_2

import "fmt"

// Клас "Автомобіль"
type Car struct {
	Speed int
}

func (c *Car) Move() {
	fmt.Println("The car is driving")
}

func (c *Car) Stop() {
	fmt.Println("The car stopped")
}

func (c *Car) ChangeSpeed(speed int) {
	c.Speed = speed
	fmt.Printf("The speed of the car has changed by %d km/h\n", speed)
}

func (c *Car) PickupPassengers(numPassengers int) {
	fmt.Printf("The car has accepted %d passengers\n", numPassengers)
}

func (c *Car) DropOffPassengers(numPassengers int) {
	fmt.Printf("The car dropped off %d passengers\n", numPassengers)
}
