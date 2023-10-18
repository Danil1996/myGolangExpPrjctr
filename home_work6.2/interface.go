package home_work6_2

// Інтерфейс "Транспортний засіб"
type Transport interface {
	Move()
	Stop()
	ChangeSpeed(speed int)
}

// Інтерфейс "Транспортний засіб з пасажирами"
type PassengerTransport interface {
	Transport
	PickupPassengers(numPassengers int)
	DropOffPassengers(numPassengers int)
}
