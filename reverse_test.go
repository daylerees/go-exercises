package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseText(t *testing.T) {
	tests := []struct{
		input string
		output string
	}{
		{"",""},
		{"a","a"},
		{"foo","oof"},
		{"bar","rab"},
		{"test","tset"},
		{"Wonderful","lufrednoW"},
	}

	for _, test := range tests {
		result := ReverseText(test.input)
		assert.Equal(t, test.output, result, "string is not reversed correctly")
	}
}
