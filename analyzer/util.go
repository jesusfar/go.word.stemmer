package analyzer

import (
	"log"
	"regexp"
	"strings"
)

func removeSpecialChar(text string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9_ ]+")

	if err != nil {
		log.Fatal(err)
		return text, err
	}

	text = reg.ReplaceAllString(text, "")

	return text, nil
}

func isStopWord(word string) bool {
	word = strings.ToLower(word)
	switch word {
	case "a", "the", "and", "of", "in", "be", "also", "as":
		return true
	}
	return false
}
