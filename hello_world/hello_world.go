package hello_world

import "fmt"

func main() {
	const two string = "Two"
	var first string = "First"
	second := "second"
	fmt.Printf("%8s merry geese lived with granny.", two)
	fmt.Printf("%8s is gray, the %s is white", first, second)
	fmt.Printf("%8s cheerful geese!", two)
	fmt.Println("***")
	fmt.Printf("The geese washed their feet in a puddle near the ditch.")
	fmt.Printf("%8s is gray, the %s is white,", first, second)
	fmt.Printf("Hid in a ditch!")
	fmt.Println("***")
	fmt.Printf("Here the granny shouts - “Oh, the geese are gone!”")
	fmt.Printf("%8s is gray, the %s is white,", first, second)
	fmt.Printf("My geese, my geese!")
	fmt.Println("***")
	fmt.Printf("The geese came out and bowed to grandma!")
	fmt.Printf("%8s is gray, the %s is white,", first, second)
	fmt.Printf("We bowed to grandma!")
	fmt.Println("***")
}
