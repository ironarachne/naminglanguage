package main

import (
	"flag"
	"fmt"

	"github.com/ironarachne/naminglanguage"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	flag.Parse()

	name := ""

	for i := 0; i < *numberOfNames; i++ {
		name = naminglanguage.GenerateName()
		fmt.Println(name)
	}
}
