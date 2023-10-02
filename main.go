package main

import (
	arrays_lice "github.com/myGolangExpPrjctr/array_slice"
)

// TODO: just run it
func main() {
	myTexts := arrays_lice.UserText{}
	myTexts.AppendText("Hello")
	myTexts.AppendText("Hello, guys")
	myTexts.AppendText("Kon'nichiwa, guys")
	myTexts.AppendText("Kon'nichiwa, girls")
	myTexts.AppendText("Hello, girls")
	myTexts.AppendText("Zdorovenki buli, guys")
	myTexts.AppendText("Zdorovenki buli, girls")
	myTexts.ShowMyInputs()
	myTexts.FindSubstringInInputs("girls")
}
