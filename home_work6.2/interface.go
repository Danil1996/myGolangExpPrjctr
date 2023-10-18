package home_work6_2

type Transport interface {
	Move()
	Stop()
	ChangeSpeed(speed int)
}

type PassengerTransport interface {
	Transport
	PickupPassengers(numPassengers int)
	DropOffPassengers(numPassengers int)
}
