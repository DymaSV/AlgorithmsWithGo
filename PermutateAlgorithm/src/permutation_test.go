package permutation

import (
	"testing"
)

// TestPermutate calls permutation.
// Check arrayr count
func TestPermutate(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{[]string{"1", "2", "3", "4"}, 24},
		{[]string{"1", "2", "3", "4", "5"}, 120},
	}
	for _, test := range tests {
		answers := make([]string, 0)
		Permutate(&test.input, 0, len(test.input), &answers)
		if len(answers) != test.expected {
			t.Error(`Permutate {} expected, {} result,`, test.expected, len(answers))
		}
	}
}
