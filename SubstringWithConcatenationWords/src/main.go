package main

import (
	"fmt"

	"../../PermutateAlgorithm/src/permutation"
)

func main() {
	arr := []string{"1", "2", "3", "4"}
	answers := make([]string, 0)
	permutation.Permutate(&arr, 0, 4, &answers)
	fmt.Println(answers)
}
