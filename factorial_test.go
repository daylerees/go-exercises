package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tests := []struct{
		input int
		output int
	}{
		{0,1},
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{5,120},
		{6,720},
		{7,5040},
		{8,40320},
		{9,362880},
		{10,3628800},
	}

	for _, test := range tests {
		result := Factorial(test.input)
		assert.Equal(t, test.output, result, "factorial result incorrect")
	}
}