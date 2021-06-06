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
		{[]string{"1", "2", "3"}, "aa1234", []int{2}},
		{[]string{"1", "2", "3"}, "2134aa", []int{0}},
		{[]string{"1", "2", "3"}, "aa4321", []int{3}},
		{[]string{"1", "2", "3"}, "24143424", []int{}},
	}
	for _, test := range tests {
		answers := Check(&test.permutations, test.str)
		for i, v := range test.expected {
			if answers[i] != v {
				t.Error(`CheckConcatenation {} expected, {} result,`, v, answers[i])
			}
		}
	}
}
