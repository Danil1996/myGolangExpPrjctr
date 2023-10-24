package home_work6

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type Package interface {
	GetSender() string
	GetRecipient() string
	Send()
}

type DeliveryService struct{}

func (deliveryService DeliveryService) SendPackage(packageToSend Package) {
	letterType := reflect.TypeOf(Letter{})
	boxType := reflect.TypeOf(Box{})
	packageType := reflect.TypeOf(packageToSend)

	if packageType == letterType {
		fmt.Println("Nova Poshta determined that this is a letter and sends it, here's the TTN:\n")
		id := uuid.New()
		fmt.Printf("%s\n", id.String())
		packageToSend.Send()
		return
	}
	if packageType == boxType {
		fmt.Println("Nova Poshta determined that this is a box and sends it, here's the TTN:\n")
		id := uuid.New()
		fmt.Printf("%s\n", id.String())
		packageToSend.Send()
		return
	}

}
