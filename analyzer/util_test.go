package analyzer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_removeSpecialChar(t *testing.T) {
	var testCases = []struct{
		text string
		output string
	} {
		{
			"hello,%!/()&",
			"hello",
		},
		{
			"The first line of Lorem Ipsum, Lorem ipsum dolor sit amet..,",
			"The first line of Lorem Ipsum Lorem ipsum dolor sit amet",
		},
	}

	for _, testCase := range testCases {
		result, _ := removeSpecialChar(testCase.text)
		assert.Equal(t, testCase.output, result)
	}

}

func Test_isStopWord(t *testing.T)  {
	var testCases = []struct{
		word string
		output bool
	} {
		{
			"a",
			true,
		},
		{
			"the",
			true,
		},
		{
			"and",
			true,
		},
		{
			"of",
			true,
		},
		{
			"in",
			true,
		},
		{
			"be",
			true,
		},
		{
			"also",
			true,
		},
		{
			"as",
			true,
		},
		{
			"hello",
			false,
		},
	}

	for _, testCase := range testCases {
		result := isStopWord(testCase.word)
		assert.Equal(t, testCase.output, result)
	}
}
