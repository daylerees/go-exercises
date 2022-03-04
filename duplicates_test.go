package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct{
		input []int
		output []int
	}{
		{
			[]int{1, 2, 2, 3, 2, 3, 4, 5, 4, 1, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{97, 21, 22, 80, 30, 35, 80, 97},
			[]int{97, 21, 22, 80, 30, 35},
		},
	}

	for _, test := range tests {
		result := RemoveDuplicates(test.input)
		assert.Equal(t, test.output, result, "duplicates are not removed")
	}
}

