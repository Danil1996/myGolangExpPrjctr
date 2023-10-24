package home_work6

import "fmt"

type Box struct {
	Sender    string
	Recipient string
}

func (b Box) GetSender() string {
	return b.Sender
}

func (b Box) GetRecipient() string {
	return b.Recipient
}

func (b Box) Send() {
	fmt.Printf("Box sent from %s to %s.\n\n", b.Sender, b.Recipient)
}
