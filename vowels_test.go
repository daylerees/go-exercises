package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVowels(t *testing.T) {
	tests := []struct{
		input string
		output int
	}{
		{"Apple",2},
		{"Hippopotamus",5},
		{"This is a sentence",6},
		{"Rhythm",0},
	}

	for _, test := range tests {
		result := GetVowels(test.input)
		assert.Equal(t, test.output, result, "incorrect number of vowels")
	}
}
