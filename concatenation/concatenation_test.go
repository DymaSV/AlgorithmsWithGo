package concatenation

import (
	"testing"
)

// TestPermutate calls permutation.
// Check arrayr count
func TestCheckConcatenation(t *testing.T) {
	var tests = []struct {
		permutations []string
		str          string
		expected     []int
	}{
		{[]string{"1", "2", "3", "4"}, "aa1234", []int{2}},
		{[]string{"1", "2", "3", "4"}, "2434aa", []int{0}},
		{[]string{"1", "2", "3", "4"}, "aa4321", []int{2}},
		{[]string{"1", "2", "3", "4"}, "24143424", []int{}},
	}
	for _, test := range tests {
		answers := CheckConcatenation(&test.permutations, test.str)
		for i, v := range test.expected {
			if answers[i] != v {
				t.Error(`CheckConcatenation {} expected, {} result,`, v, answers[i])
			}
		}
	}
}
