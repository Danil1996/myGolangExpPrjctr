package controlFlow

import (
	"fmt"
	"regexp"
)

type Character struct {
	Name string
}

type Quest struct{}

func (quest *Quest) StartGame(character Character) {
Start:
	initializeIntro(&character)

	if containsNonEnglishCharacters(character.Name) {
		fmt.Println("!!!WARNING!!! In English please.")
		goto Start
	}

	fmt.Printf("Okay %s, let's play \n", character.Name)

	fmt.Print(" ------ FIRST PART  ------ \n")

	fmt.Printf("%s, you wake up in an unfamiliar place with some things. You don't remember anything. There are three doors in front of you. You can choose one of them. Your decision will influence the development of events:\n", character.Name)

	fmt.Printf("Choose the door. left? right? forward?")

FirstChoice:
	var firstChoice string
	fmt.Scan(&firstChoice)

	if !containsDoors(firstChoice) {
		fmt.Println("!!!WARNING!!! left? right? forward?")
		goto FirstChoice
	}

	switch firstChoice {
	case "left":
		fmt.Printf("You decide to go left. Behind the door is a corridor with many doors. Select one of them and enter the room. You see a comfortable chair, a bottle of delicious beer in the cup holder and a bucket of popcorn on the table nearby. Isn't this a sign?, since you have a bottle opener in your pocket. You sit down, open a beer and a voice outside asks what movie do you want to watch?.\n")
	LeftDoor:
		var film string
		fmt.Scan(&film)
		if containsNonEnglishCharacters(film) {
			fmt.Println("!!!WARNING!!! In English please.")
			goto LeftDoor
		}
		fmt.Printf("The film \"%s\" turned out to be amazing. You had a good time. At the very end, when you finished your beer and there was no popcorn left, you woke up in your bed. It was an amazing dream\n", film)
	case "forward":
		fmt.Printf("You decide to go straight. A long corridor with painted walls awaits you. There is light at the end of the corridor. Having reached the light, you realize that you are looking directly at \"Giaconda\". This is the Louvre. After walking around the Louvre for several hours, you got tired and sat down to rest. We fell asleep and woke up in our bed with an amazing mood.\n")
	case "right":
		fmt.Printf("You decide to go right. Outside the door, go out into the street, where you see an unfamiliar city. And you have several options on what to go with, let me list them.\n")

		transports := [3]string{"car", "bike", "motorbike"}
		for _, transports := range transports {
			fmt.Println(transports + "\n")
		}
		fmt.Printf("What are we going to ride?\n")
	RightDoor:
		var transport string
		fmt.Scan(&transport)
		if !containsTransport(transport) {
			fmt.Println("!!!WARNING!!! Car? Bike? Motorbike?")
			goto RightDoor
		}
		fmt.Printf("Having taken a ride on %s, you enjoyed the views of a city unknown to you and after several hours of travel you were tired and your eyes began to close on their own. At some point you felt suspended and immediately woke up at home.\n", transport)
	}

	fmt.Println(" ------ THE END ------ \n")

	fmt.Printf("%s. Don’t be upset, this quest can be developed further, this requires time and homework that could be associated with this))) Thank you for your attention, see you soon...", character.Name)

	fmt.Printf("Want to try again? yes? no?")
End:
	var choice string
	fmt.Scan(&choice)
	if choice == "yes" {
		goto Start
	}
	if choice == "no" {
		fmt.Printf("As you wish")
	} else {
		fmt.Printf("CAN YOU BE A NORMAL PERSON AND RESPECT A MACHINE? Choose either yes or no?")
		goto End
	}

}

func initializeIntro(character *Character) {
	fmt.Println("Welcome to the quest: \"New World\"")
	fmt.Println("What is your name?")
	fmt.Scan(&character.Name)
}

func containsNonEnglishCharacters(input string) bool {
	nonEnglishRegex := regexp.MustCompile(`[а-яА-ЯґҐєЄіІїЇёЁ0-9]`)
	return nonEnglishRegex.MatchString(input)
}

func containsDoors(input string) bool {
	switch input {
	case "right":
		return true
	case "left":
		return true
	case "forward":
		return true
	default:
		return false
	}
}

func containsTransport(input string) bool {
	switch input {
	case "car":
		return true
	case "bike":
		return true
	case "motorbike":
		return true
	default:
		return false
	}
}
