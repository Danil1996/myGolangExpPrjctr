package main

import (
	"fmt"

	"github.com/Danil1996/myGolangExpPrjctr/home_work6.1"
	home_work6_2 "github.com/Danil1996/myGolangExpPrjctr/home_work6.2"
)

func firstTask() {
	package1 := home_work6.Box{Sender: "Daniil", Recipient: "Greg"}
	package2 := home_work6.Letter{Sender: "Max", Recipient: "Viktor"}

	NovaPoshta := home_work6.DeliveryService{}

	NovaPoshta.SendPackage(package1)
	NovaPoshta.SendPackage(package2)
}

func secondTask() {
	route := &home_work6_2.Route{}

	fmt.Println("- My journey began with the fact that I got into my car and went to the airport.I'm heading from Odessa to Dresden")

	car := &home_work6_2.Car{}
	route.AddTransport(car)
	car.PickupPassengers(1)
	car.Move()
	car.ChangeSpeed(120)

	fmt.Println("- I got to the airport.")

	car.ChangeSpeed(5)
	car.Stop()
	car.DropOffPassengers(1)

	fmt.Println("- After I got to the airport, I parked my car in the parking lot and went to board the plane.")

	plane := &home_work6_2.Plane{}

	fmt.Println("- After going through check-in and passport control, I was already on the plane along with 150 other people.")

	route.AddTransport(plane)
	plane.PickupPassengers(151)
	plane.Move()
	plane.ChangeSpeed(900)

	fmt.Println("- Time has flown by quite quickly, and now we are landing in Berlin")
	plane.ChangeSpeed(50)
	plane.Stop()
	plane.DropOffPassengers(151)

	train := &home_work6_2.Train{}

	fmt.Println("- There's one last push left. Now I'm already at the train station and getting ready to board the train.")

	route.AddTransport(train)
	train.PickupPassengers(300)
	train.Move()
	train.ChangeSpeed(150)

	fmt.Println("- Several hours passed and now I'm in Dresden")
	train.ChangeSpeed(10)
	train.Stop()
	train.DropOffPassengers(300)

	fmt.Println("- It was an exciting journey, in just a few hours I drove and flew almost 1500 km using three flights")
	route.ShowTransportList()

}

func main() {
	fmt.Println("!!!! FIRST TASK START!!!!\n")
	firstTask()
	fmt.Println("!!!! FIRST TASK END !!!!\n")

	fmt.Println("!!!! SECOND TASK START!!!!\n")
	secondTask()
	fmt.Println("!!!! SECOND TASK END !!!!\n")
}
