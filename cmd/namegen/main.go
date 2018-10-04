package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/ironarachne/naminglanguage"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	typeOfName := flag.String("type", "person", "Type of name to generate (person or place)")
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")
	firstNameOnly := flag.Bool("f",false,"Generate first name")
	lastNameOnly := flag.Bool("l",false,"Generate last name")
	meaning := flag.Bool("m",false,"Generate name meaning for first and/or last name")

	flag.Parse()

	name := ""

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	for i := 0; i < *numberOfNames; i++ {
		if *typeOfName == "person" {
			name = naminglanguage.GeneratePersonName(*firstNameOnly,*lastNameOnly,*meaning)
		} else {
			name = naminglanguage.GeneratePlaceName()
		}

		fmt.Println(name)
	}
}
