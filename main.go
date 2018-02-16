package main

import (
	"os"
	"fmt"
	"github.com/jesusfar/go.word.stemmer/analyzer"
	"encoding/json"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	switch args[0] {
	case "analyze":
		if len(args) >= 2 {
			text := args[1]
			result := analyzer.Analyze(text)
			resultMarshal, _ := json.MarshalIndent(result, "", "\t")
			fmt.Println(string(resultMarshal))
		}
	default:
		printHelp()

	}
}
func printHelp() {
	fmt.Println(`
go.word.stemmer is a tool for analyze a paragraph and returns a stemmed word analysis.

Usage: go.word.stemmer <command>

Commands:
  analyze            Analyze a paragraph

	`)
}