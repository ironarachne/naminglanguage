package naminglanguage

import (
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

// GeneratePersonName generates a random name for a person
func GeneratePersonName() string {
	language := GenerateLanguage()
	firstName := randomWord(language)
	firstPartOfLastName := randomWord(language)
	secondPartOfLastName := randomWord(language)
	name := firstName.Word + " " + firstPartOfLastName.Word + secondPartOfLastName.Word
	name = strings.Title(name)
	return name
}

// GeneratePlaceName generates a random name for a place
func GeneratePlaceName() string {
	language := GenerateLanguage()
	firstPart := randomWord(language)
	secondPart := randomWord(language)
	name := firstPart.Word + secondPart.Word
	name = strings.Title(name)
	return name
}

// GenerateLanguage generates a language
func GenerateLanguage() Language {
	syllables := generateSyllables()
	words := generateWords(syllables)
	wordOrder := randomWordOrder()
	language := Language{words, wordOrder}

	return language
}
