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
func GeneratePersonName(first bool, last bool, meaning bool) string {
	language := GenerateLanguage()
	firstName := randomNoun(language)
	firstPartOfLastName := randomVerb(language)
	secondPartOfLastName := randomNoun(language)
	name := ""
	// First Name Only
	if (first) {
		name += firstName.Word + " "
	}
	// Last Name Only
	if (last) {
		name += firstPartOfLastName.Word + secondPartOfLastName.Word + " "
	}
	// First and Land name with no meaning
	if (meaning) {
		if(first && last){
		name +=	"(" + firstName.Meaning + ", " + nounFromVerb(firstPartOfLastName.Meaning) + " of " + pluralizeNoun(secondPartOfLastName.Meaning) + ")"
		} else if (first) {
			name += "(" + firstName.Meaning + ")"
		} else if (last){
			name +=	"(" + nounFromVerb(firstPartOfLastName.Meaning) + " of " + pluralizeNoun(secondPartOfLastName.Meaning) + ")"
		}
	}
	// Defualt of Full name + Meaning
	if (!first && !last) {
		name += firstName.Word + " " + firstPartOfLastName.Word + secondPartOfLastName.Word + " (" + firstName.Meaning + ", " + nounFromVerb(firstPartOfLastName.Meaning) + " of " + pluralizeNoun(secondPartOfLastName.Meaning) + ")"
	}

	name = strings.Title(name)
	return name
}

// GeneratePlaceName generates a random name for a place
func GeneratePlaceName() string {
	language := GenerateLanguage()
	firstPart := randomAdjective(language)
	secondPart := randomNoun(language)
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
