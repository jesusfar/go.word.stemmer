package analyzer

import (
	"log"
	"strings"
	"github.com/jesusfar/snowball"
	"sort"
)

type Result struct {
	Word             string `json:"word"`
	TotalOccurrences int `json:"total-occurrences"`
	SentenceIndexes []int `json:"sentence-indexes"`
}

type AnalysisOutPut struct {
	Results []Result `json:"results"`
}

type Analyzer interface {
	Analyze(paragraph string) AnalysisOutPut
}

//Analyze return a list of stemmed word analysis
func Analyze(paragraph string) AnalysisOutPut {
	var results []Result
	var uniqueWords []string
	stemmedWords := make(map[string]Result)

	paragraph, err := removeSpecialChar(paragraph)

	if err != nil {
		log.Fatal(err)
	}

	words := strings.Fields(paragraph)

	for index, word := range words {
		//Ignore Stop Words
		if isStopWord(word) {
			continue
		}
		stemmed, err := snowball.Stem(word, "english", true)
		if err == nil {
			value, ok := stemmedWords[word]
			if ok {
				value.Word = stemmed
				value.TotalOccurrences++
				value.SentenceIndexes = append(value.SentenceIndexes, index)
				stemmedWords[word] = value
			} else {
				stemmedWords[word] = Result{
					Word: stemmed,
					TotalOccurrences: 1,
					SentenceIndexes: []int{index},
				}
				uniqueWords = append(uniqueWords, word)
			}

		}

	}

	// Sort key words
	sort.Strings(uniqueWords)

	// Prepare result
	for _, uniqueWord := range uniqueWords {
		stemmedWord, ok := stemmedWords[uniqueWord]
		if ok {
			results = append(results, stemmedWord)
		}
	}

	return AnalysisOutPut{ Results: results }
}
