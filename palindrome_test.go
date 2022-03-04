package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct{
		input string
		output bool
	}{
		{"a",true},
		{"level",true},
		{"madam",true},
		{"radar",true},
		{"redder",true},
		{"repaper",true},
		{"sagas",true},
		{"solos",true},
		{"anna",true},
		{"kayak",true},
		{"apple",false},
		{"battery",false},
		{"car",false},
		{"dog",false},
		{"keyboard",false},
		{"pencil",false},
		{"phone",false},
		{"problem",false},
		{"transport",false},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		assert.Equal(t, test.output, result, "palindrome determination incorrect")
	}
}

