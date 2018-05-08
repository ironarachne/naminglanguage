package naminglanguage

import (
	"math/rand"
	"strings"
	"time"
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

// GenerateName generates a random name
func GenerateName() string {
	language := GenerateLanguage()
	firstName := randomWord(language)
	firstPartOfLastName := randomWord(language)
	secondPartOfLastName := randomWord(language)
	name := firstName.Word + " " + firstPartOfLastName.Word + secondPartOfLastName.Word
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

func generateSyllable(consonants string, vowels string, order []string) string {
	syllable := ""
	component := ""
	for _, orderItem := range order {
		component = randomCharacterFromString(consonants)
		if orderItem == "V" {
			component = randomCharacterFromString(vowels)
		}
		syllable += component
	}
	return syllable
}

func generateSyllables() []string {
	syllable := ""
	consonants := randomConsonantSet()
	vowels := randomVowelSet()
	order := randomSyllableOrder()
	var syllables []string

	for i := 0; i < 10; i++ {
		syllable = generateSyllable(consonants, vowels, order)
		syllables = append(syllables, syllable)
	}

	return syllables
}

func generateWords(syllables []string) []Word {
	var words []Word
	word := Word{}
	wordString := ""
	wordSyllableLength := randomWordSyllableLength()
	for i := 0; i < 50; i++ {
		wordString = ""
		for j := 0; j < wordSyllableLength; j++ {
			wordString += randomSyllable(syllables)
		}
		word = Word{wordString, "", ""}
		words = append(words, word)
	}
	return words
}

func randomConsonantSet() string {
	rand.Seed(time.Now().UnixNano())
	consonantSets := [1]string{"ptkmnsl"}
	return consonantSets[rand.Intn(len(consonantSets))]
}

func randomCharacterFromString(items string) string {
	rand.Seed(time.Now().UnixNano())
	characters := []byte(items)
	return string(characters[rand.Intn(len(characters))])
}

func randomVowelSet() string {
	rand.Seed(time.Now().UnixNano())
	vowelSets := [1]string{"aeiou"}
	return vowelSets[rand.Intn(len(vowelSets))]
}

func randomWord(language Language) Word {
	rand.Seed(time.Now().UnixNano())
	return language.Dictionary[rand.Intn(len(language.Dictionary))]
}

func randomSyllable(syllables []string) string {
	rand.Seed(time.Now().UnixNano())
	return syllables[rand.Intn(len(syllables))]
}

func randomSyllableOrder() []string {
	rand.Seed(time.Now().UnixNano())
	syllableOrders := [1][]string{{"C", "V", "C"}}
	return syllableOrders[rand.Intn(len(syllableOrders))]
}

func randomWordOrder() []string {
	rand.Seed(time.Now().UnixNano())
	wordOrders := [1][]string{{"S", "V", "O"}}
	return wordOrders[rand.Intn(len(wordOrders))]
}

func randomWordSyllableLength() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) + 1
}
