package home_work6_2

import "fmt"

// Клас "Потяг"
type Train struct {
	Speed int
}

func (t *Train) Move() {
	fmt.Println("The train is going")
}

func (t *Train) Stop() {
	fmt.Println("The train stopped")
}

func (t *Train) ChangeSpeed(speed int) {
	t.Speed = speed
	fmt.Printf("The speed of the train has changed by %d km/h\n", speed)
}

func (t *Train) PickupPassengers(numPassengers int) {
	fmt.Printf("The train has received %d passengers\n", numPassengers)
}

func (t *Train) DropOffPassengers(numPassengers int) {
	fmt.Printf("The train dropped off %d passengers\n", numPassengers)
}
