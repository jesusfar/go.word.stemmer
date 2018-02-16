# Word Analyzer stemmer PoC

This is a basic PoC in order to test the snowball library for natural language processing.

## How to install

```
$ go get github.com/jesusfar/go.word.stemmer
```

## Running test
Go to the project
```
$ cd $GOPATH/src/github.com/jesusfar/go.word.stemmer
``` 
Run unit tests
``` 
$ go test -v ./...
```
Run benchmark tests
```
$ go test -v ./... -bench .
```

## Example usage:

```
go.word.stemmer analyze "Take this paragraph of text and return an alphabetized list of ALL unique words."
```

## Licence
MIT licence.
