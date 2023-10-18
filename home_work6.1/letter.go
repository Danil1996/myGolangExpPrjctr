package home_work6

import "fmt"

// Створення структури "Конверт", яка реалізує інтерфейс "Посилка"
type Letter struct {
	Sender    string
	Recipient string
}

func (e Letter) GetSender() string {
	return e.Sender
}

func (e Letter) GetRecipient() string {
	return e.Recipient
}

func (e Letter) Send() {
	fmt.Printf("Letter send from %s to %s.\n", e.Sender, e.Recipient)
}
