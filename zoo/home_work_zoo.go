package zoo

import "fmt"

type animal struct {
	kind     string
	name     string
	skill    string
	quantity int
}

type staff struct {
	name     string
	position string
	property string
}

type zoo struct {
	isCageExist      bool
	animalLivingArea string
}

type place struct {
	name string
	size string
	kind zoo
}

type ZooStoryManager struct{}

func (staff *staff) defineStaffPropertyByPosition() {
	if staff.position == "director" {
		staff.property = "knows every animal by name"
	}
	if staff.position == "manager" {
		staff.property = "expert in animal care and always found ways to improve living conditions"
	}
}

func (animal *animal) defineAnimalSkillByKind() {
	if animal.kind == "cheetah" {
		animal.skill = "running skills"
	}
	if animal.kind == "turtle" {
		animal.skill = "crawl slowly"
	}
	if animal.kind == "flamingo" {
		animal.skill = "swim"
	}
	if animal.kind == "monkey" {
		animal.skill = "dancing merrily"
	}
}

func (animal *animal) isAnimalInOneQuantity() {
	if animal.quantity > 1 {
		animal.name = animal.kind + "s"
	}
}

func (zooStoryManager *ZooStoryManager) composeStory(zoo place, director staff, manager staff, cheetah animal, turtle animal, flamingo animal, monkey animal) string {

	var firstPart string = fmt.Sprintf("In a small town surrounded by picturesque forests, there was a %s but charming zoo called \"%s\". This zoo was home to many amazing animals and attracted visitors from far away places.The \"%s\" zoo was known for its unique concept. Are there traditional cells there? %t, it provided them with more natural and spacious habitats such as %s. This made the zoo especially popular among animal lovers.", zoo.size, zoo.name, zoo.name, zoo.kind.isCageExist, zoo.kind.animalLivingArea)

	var secondPart string = fmt.Sprintf("The zoo staff were passionate and caring. At the head of the zoo was the %s %s, who was always in touch with all the animals and %s. Her right hand, %s %s, was %s.", director.position, director.name, director.property, manager.position, manager.name, manager.property)

	var thirdPart string = fmt.Sprintf(" Among the zoo's most popular residents was %s, a young %s who always delighted visitors with her lightning-fast %s. Closer to the entrance there was an area where playful %s lived, entertaining the children and %s.", cheetah.name, cheetah.kind, cheetah.skill, monkey.name, monkey.skill)

	var fourthPart string = fmt.Sprintf("But the most amazing thing was the open pond where %s and %s lived. Visitors could watch %s %s in the water, creating stunning images, while %s %s along the shore.", flamingo.name, turtle.name, flamingo.name, flamingo.skill, turtle.name, turtle.skill)

	var fifthPart string = fmt.Sprintln("In winter, the zoo held festivals with candles, which gave the place a special magical feel.The Forest Wonders Zoo was not only a place of entertainment, but also a place where one could learn about the richness of the animal world and observe its inhabitants in more natural conditions. This story of the zoo, its animals and its caring staff always delights visitors and reminds them of the importance of wildlife conservation.")

	return firstPart + secondPart + thirdPart + fourthPart + fifthPart
}

func (zooStoryManager *ZooStoryManager) TellStory() string {
	zoo := place{
		name: "Forest Wonders",
		size: "small",
		kind: zoo{
			isCageExist:      false,
			animalLivingArea: "open and fenced space",
		},
	}

	director := staff{
		name:     "Margaret",
		position: "director",
	}
	director.defineStaffPropertyByPosition()

	manager := staff{
		name:     "Tom",
		position: "manager",
	}
	manager.defineStaffPropertyByPosition()

	cheetah := animal{
		kind:     "cheetah",
		name:     "Alice",
		quantity: 1,
	}
	cheetah.defineAnimalSkillByKind()
	cheetah.isAnimalInOneQuantity()

	turtle := animal{
		kind:     "turtle",
		quantity: 10,
	}
	turtle.defineAnimalSkillByKind()
	turtle.isAnimalInOneQuantity()

	flamingo := animal{
		kind:     "flamingo",
		quantity: 14,
	}
	flamingo.defineAnimalSkillByKind()
	flamingo.isAnimalInOneQuantity()

	monkey := animal{
		kind:     "monkey",
		quantity: 9,
	}
	monkey.defineAnimalSkillByKind()
	monkey.isAnimalInOneQuantity()

	return zooStoryManager.composeStory(zoo, director, manager, cheetah, turtle, flamingo, monkey)
}
