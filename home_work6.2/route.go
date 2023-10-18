package home_work6_2

import "fmt"

// Клас "Маршрут"
type Route struct {
	TransportList []Transport
}

func (r *Route) AddTransport(transport Transport) {
	r.TransportList = append(r.TransportList, transport)
}

func (r *Route) ShowTransportList() {
	fmt.Println("Vehicles were on the route:")
	for _, transport := range r.TransportList {
		fmt.Printf("%T\n", transport)
	}
}
