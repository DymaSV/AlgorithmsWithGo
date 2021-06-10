package concatenation

import (
	"testing"
)

// TestPermutate calls permutation.
// Check arrayr count
func TestCheckConcatenation(t *testing.T) {
	var tests = []struct {
		permutation string
		str         string
		expected    []int
	}{
		{"123", "2a123b321", []int{2}},
		{"123", "21123a", []int{2}},
		{"123", "a12321", []int{1}},
		{"123", "24143424", []int{}},
	}
	for _, test := range tests {
		answers := Check(test.permutation, test.str)
		for i, v := range test.expected {
			if answers[i] != v {
				t.Error(`CheckConcatenation {} expected, {} result,`, v, answers[i])
			}
		}
	}
}
