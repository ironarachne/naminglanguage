package main

import (
	"fmt"

	"github.com/ironarachne/naminglanguage"
)

func main() {
	language := naminglanguage.GenerateLanguage()

	for _, word := range language.Dictionary {
		fmt.Println(word.Word + ": " + word.Meaning)
	}
}
