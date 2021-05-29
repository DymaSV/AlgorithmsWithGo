package concatenation

import (
	"strings"
)

func CheckConcatenation(variants *[]string, str string) []int {
	arr := strings.Split(str, "")
	index := 0
	results := make([]int, 0)
	resIndex := 0
	for _, variant := range *variants {
		for _, v := range strings.Split(variant, "") {
			if arr[index] == v {
				results[resIndex] = index
				index = index + 1
				if index == len(str) {
					return results
				}
			} else {
				index = 0
			}
		}
	}
	return []int{2}
}
