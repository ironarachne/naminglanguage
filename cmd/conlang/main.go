package main

import (
	"fmt"
	"flag"
	"time"
	"math/rand"

	"github.com/ironarachne/naminglanguage"
)

func main() {
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")
	flag.Parse()

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	language := naminglanguage.GenerateLanguage()

	for _, word := range language.Dictionary {
		fmt.Println(word.Word + ": " + word.Meaning)
	}
}

