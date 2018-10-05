package naminglanguage

import (
	"fmt"
	"strings"
)

// Word is a word
type Word struct {
	Word    string
	Meaning string
	Part    string
}

// Language contains a dictionary of words and a word order
type Language struct {
	Dictionary []Word
	WordOrder  []string
}

// Person is a representation of a person.
type Person struct {
	FirstName        string `json:",omitempty"`
	LastName         string `json:",omitempty"`
	FirstNameMeaning string `json:",omitempty"`
	LastNameMeaning  string `json:",omitempty"`
}

// String returns the string version of Person
func (p *Person) String() string {
	var fields []string

	if p.FirstName != "" {
		fields = append(fields, p.FirstName)
	}

	if p.LastName != "" {
		fields = append(fields, p.LastName)
	}

	if p.FirstNameMeaning != "" {
		fields = append(fields, fmt.Sprintf("(%s, %s)", p.FirstNameMeaning, p.LastNameMeaning))
	}

	return strings.Join(fields, " ")
}

// Place is representation of a place.
type Place struct {
	Name string
	// FIXME: unsure how to implement this, it was in the original issue for json though.
	// Meaning string
}

// GeneratePerson generates a random name for a person.
func GeneratePerson() *Person {
	language := GenerateLanguage()
	firstName := randomNoun(language)
	firstPartOfLastName := randomVerb(language)
	secondPartOfLastName := randomNoun(language)

	return &Person{
		FirstName:        strings.Title(firstName.Word),
		LastName:         strings.Title(firstPartOfLastName.Word + secondPartOfLastName.Word),
		FirstNameMeaning: strings.Title(firstName.Meaning),
		LastNameMeaning:  strings.Title(nounFromVerb(firstPartOfLastName.Meaning) + " of " + pluralizeNoun(secondPartOfLastName.Meaning)),
	}
}

// GeneratePlace generates a random name for a place.
func GeneratePlace() *Place {
	language := GenerateLanguage()
	firstPart := randomAdjective(language)
	secondPart := randomNoun(language)

	return &Place{
		Name: firstPart.Word + secondPart.Word,
	}
}

// GenerateLanguage generates a language
func GenerateLanguage() Language {
	syllables := generateSyllables()
	words := generateWords(syllables)
	wordOrder := randomWordOrder()
	language := Language{words, wordOrder}

	return language
}

// TODO:
// 	GeneratePersonFromLanguage(lang *Language) *Person
//	GeneratePlaceFromLanguage(lang *Language) *Place
