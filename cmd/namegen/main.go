package main

import (
	"flag"
	"fmt"

	"github.com/ironarachne/naminglanguage"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	typeOfName := flag.String("type", "person", "Type of name to generate (person or place)")
	flag.Parse()

	name := ""

	for i := 0; i < *numberOfNames; i++ {
		if *typeOfName == "person" {
			name = naminglanguage.GeneratePersonName()
		} else {
			name = naminglanguage.GeneratePlaceName()
		}

		fmt.Println(name)
	}
}
