package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/random"
)

func main() {
	randomSeed := flag.String("s", "now", "Optional random generator seed")
	flag.Parse()

	if *randomSeed == "now" {
		rand.Seed(time.Now().UnixNano())
	} else {
		random.SeedFromString(*randomSeed)
	}

	language := naminglanguage.GenerateLanguage()

	for _, word := range language.Dictionary {
		fmt.Println(word.Word + ": " + word.Meaning)
	}
}
