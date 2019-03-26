package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/random"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	typeOfName := flag.String("type", "person", "Type of name to generate (person or place)")
	randomSeed := flag.String("s", "now", "Optional random generator seed")
	firstNameOnly := flag.Bool("f", false, "Generate first name")
	lastNameOnly := flag.Bool("l", false, "Generate last name")
	meaning := flag.Bool("m", false, "Generate name meaning for first and/or last name")
	asJSON := flag.Bool("json", false, "return values as json")
	flag.Parse()

	if *randomSeed == "now" {
		rand.Seed(time.Now().UnixNano())
	} else {
		random.SeedFromString(*randomSeed)
	}

	var err error
	if *asJSON {
		err = _asJSON(*numberOfNames, *typeOfName, *firstNameOnly, *lastNameOnly, *meaning)
	} else {
		_asText(*numberOfNames, *typeOfName, *firstNameOnly, *lastNameOnly, *meaning)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func _asJSON(c int, t string, firstNameOnly, lastNameOnly, meaning bool) error {
	final := make([]interface{}, 0, c)

	for i := 0; i < c; i++ {
		if t == "person" {
			p := _genPerson(firstNameOnly, lastNameOnly, meaning)
			final = append(final, p)
		} else {
			p := naminglanguage.GeneratePlace()
			final = append(final, p)
		}
	}

	b, err := json.MarshalIndent(final, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("%s", b)
	return nil
}

func _asText(c int, t string, firstNameOnly, lastNameOnly, meaning bool) {
	var name string
	for i := 0; i < c; i++ {
		if t == "person" {
			name = _genPerson(firstNameOnly, lastNameOnly, meaning).String()
		} else {
			name = naminglanguage.GeneratePlace().Name
		}

		fmt.Println(name)
	}
}

func _genPerson(firstNameOnly, lastNameOnly, meaning bool) *naminglanguage.Person {
	p := naminglanguage.GeneratePerson()

	switch {
	case firstNameOnly:
		p.LastName = ""
		p.FirstNameMeaning = ""
		p.LastNameMeaning = ""
	case lastNameOnly:
		p.FirstName = ""
		p.FirstNameMeaning = ""
		p.LastNameMeaning = ""
	default:
		if !meaning {
			p.FirstNameMeaning = ""
			p.LastNameMeaning = ""
		}
	}

	return p
}
