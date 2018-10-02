package naminglanguage

import (
	"math/rand"
	"time"
)

func isUnique(word string, wordList []Word) bool {
	for _, entry := range wordList {
		if word == entry.Word {
			return false
		}
	}
	return true
}

func nounFromVerb(word string) string {
	lastCharacter := string(word[len(word)-1:])
	if (lastCharacter == "e") {
		return word + "r"
	}
	return word + "er"
}

func pluralizeNoun(word string) string {
	lastCharacter := string(word[len(word)-1:])
	if (lastCharacter == "s") {
		return word + "ses"
	}
	return word + "s"
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

	for i := 0; i < 150; i++ {
		syllable = generateSyllable(consonants, vowels, sibilants, glides, finals, order)
		for stringInSlice(syllable, badCombinations) {
			syllable = generateSyllable(consonants, vowels, sibilants, glides, finals, order)
		}
		syllables = append(syllables, syllable)
	}

	return syllables
}

func generateWordsForType(wordType string, wordList []string, existingWords []Word, maxSyllables int, syllables []string) []Word {
	var words []Word
	word := Word{}
	wordString := ""
	unique := false

	for _, wordMeaning := range wordList {
		unique = false
		for !unique {
			wordString = ""
			for j := 0; j < maxSyllables; j++ {
				wordString += randomSyllable(syllables)
			}
			if isUnique(wordString, existingWords) {
				unique = true
			}
		}
		word = Word{wordString, wordMeaning, wordType}
		words = append(words, word)
	}

	return words
}

func generateWords(syllables []string) []Word {
	var words []Word

	words = append(words, generateWordsForType("article", articles, words, 1, syllables)...)
	words = append(words, generateWordsForType("adjective", adjectives, words, randomWordSyllableLength(2), syllables)...)
	words = append(words, generateWordsForType("adverb", adverbs, words, randomWordSyllableLength(2), syllables)...)
	words = append(words, generateWordsForType("noun", nouns, words, randomWordSyllableLength(3), syllables)...)
	words = append(words, generateWordsForType("verb", verbs, words, randomWordSyllableLength(3), syllables)...)
	words = append(words, generateWordsForType("conjunction", conjunctions, words, 1, syllables)...)
	words = append(words, generateWordsForType("pronoun", pronouns, words, 1, syllables)...)

	return words
}

func getWordsByType(language Language, wordType string) []Word {
	var words []Word
	for _, word := range language.Dictionary {
		if word.Part == wordType {
			words = append(words, word)
		}
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

func randomAdjective(language Language) Word {
	rand.Seed(time.Now().UnixNano())
	words := getWordsByType(language, "adjective")
	return words[rand.Intn(len(words))]
}

func randomNoun(language Language) Word {
	rand.Seed(time.Now().UnixNano())
	words := getWordsByType(language, "noun")
	return words[rand.Intn(len(words))]
}

func randomVerb(language Language) Word {
	rand.Seed(time.Now().UnixNano())
	words := getWordsByType(language, "verb")
	return words[rand.Intn(len(words))]
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
	syllableOrders := [][]string{{"C", "V", "C"}, {"S?", "C", "V", "C"}, {"S?", "C", "V", "F"}, {"C", "V"}}
	return syllableOrders[rand.Intn(len(syllableOrders))]
}

func randomWordOrder() []string {
	rand.Seed(time.Now().UnixNano())
	wordOrders := [][]string{{"S", "V", "O"}}
	return wordOrders[rand.Intn(len(wordOrders))]
}

func randomWordSyllableLength(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1
}
