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

func generateSyllable(consonants string, vowels string, sibilants string, glides string, finals string, order []string) string {
	syllable := ""
	component := ""
	for _, orderItem := range order {
		component = randomCharacterFromString(consonants)
		if orderItem == "V" {
			component = randomCharacterFromString(vowels)
		} else if orderItem == "S?" {
			component = optionalCharacter(randomCharacterFromString(sibilants))
		} else if orderItem == "G" {
			component = randomCharacterFromString(glides)
		} else if orderItem == "F" {
			component = randomCharacterFromString(finals)
		}
		syllable += component
	}
	return syllable
}

func generateSyllables() []string {
	syllable := ""
	consonants := randomConsonantSet()
	vowels := randomVowelSet()
	sibilants := randomSibilantSet()
	glides := randomGlideSet()
	finals := randomFinalSet()
	order := randomSyllableOrder()
	var syllables []string
	badCombinations := []string{"sd", "bx", "cx", "zd"}
	doubleLetter := "aa"
	for _, letter := range "abcdefghijklmnopqrstuvwyxz" {
		doubleLetter = string(letter + letter)
		badCombinations = append(badCombinations, doubleLetter)
	}

	for i := 0; i < 10; i++ {
		syllable = generateSyllable(consonants, vowels, sibilants, glides, finals, order)
		for stringInSlice(syllable, badCombinations) {
			syllable = generateSyllable(consonants, vowels, sibilants, glides, finals, order)
		}
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

func optionalCharacter(character string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	if i >= 6 {
		return character
	}
	return ""
}

func randomConsonantSet() string {
	rand.Seed(time.Now().UnixNano())
	consonantSets := []string{"ptkmnsl", "ptkbdgmnlrsz", "ptkqvsgrmn", "tkdgmns", "tksdbqxmnlrwj"}
	return consonantSets[rand.Intn(len(consonantSets))]
}

func randomCharacterFromString(items string) string {
	rand.Seed(time.Now().UnixNano())
	characters := []byte(items)
	return string(characters[rand.Intn(len(characters))])
}

func randomFinalSet() string {
	rand.Seed(time.Now().UnixNano())
	finalSets := []string{"mn", "sk", "sz"}
	return finalSets[rand.Intn(len(finalSets))]
}

func randomGlideSet() string {
	rand.Seed(time.Now().UnixNano())
	glideSets := []string{"l", "r", "lr", "lrw", "lrwj"}
	return glideSets[rand.Intn(len(glideSets))]
}

func randomSibilantSet() string {
	rand.Seed(time.Now().UnixNano())
	sibilantSets := []string{"s", "sf"}
	return sibilantSets[rand.Intn(len(sibilantSets))]
}

func randomVowelSet() string {
	rand.Seed(time.Now().UnixNano())
	vowelSets := []string{"aeiou", "aiu"}
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
	syllableOrders := [][]string{{"C", "V", "C"}, {"S?", "C", "V", "C"}, {"S?", "C", "V", "F"}}
	return syllableOrders[rand.Intn(len(syllableOrders))]
}

func randomWordOrder() []string {
	rand.Seed(time.Now().UnixNano())
	wordOrders := [][]string{{"S", "V", "O"}}
	return wordOrders[rand.Intn(len(wordOrders))]
}

func randomWordSyllableLength() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) + 1
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
