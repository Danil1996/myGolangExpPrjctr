package arrays_lice

import (
	"fmt"
	"regexp"
	"strings"
)

type UserText struct {
	Inputs []string
}

func (userText *UserText) AppendText(text string) {
	if containsNonEnglishCharacters(text) {
		fmt.Println("!!!WARNING!!! In English please.")
		return
	}
	userText.Inputs = append(userText.Inputs, text)

}

func (userText *UserText) ShowMyInputs() {
	for _, text := range userText.Inputs {
		fmt.Println(text)
	}
}

func (userText *UserText) FindSubstringInInputs(substring string) {
	var sliceOfMatches []string
	for _, string := range userText.Inputs {
		if strings.Contains(string, substring) {
			sliceOfMatches = append(sliceOfMatches, string)
		}
	}
	fmt.Printf("%v\n", sliceOfMatches)
}

func containsNonEnglishCharacters(input string) bool {
	nonEnglishRegex := regexp.MustCompile(`[а-яА-ЯґҐєЄіІїЇёЁ]`)
	return nonEnglishRegex.MatchString(input)
}
