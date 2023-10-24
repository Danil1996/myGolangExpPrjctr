package home_work6_2

import "fmt"

type Plane struct {
	Speed int
}

func (p *Plane) Move() {
	fmt.Println("The plane is flying")
}

func (p *Plane) Stop() {
	fmt.Println("Plane landed")
}

func (p *Plane) ChangeSpeed(speed int) {
	p.Speed = speed
	fmt.Printf("The speed of the plane has changed by %d km/h\n", speed)
}

func (p *Plane) PickupPassengers(numPassengers int) {
	fmt.Printf("The plane has accepted %d passengers\n", numPassengers)
}

func (p *Plane) DropOffPassengers(numPassengers int) {
	fmt.Printf("The plane dropped off %d passengers\n", numPassengers)
}
