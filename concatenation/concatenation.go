package concatenation

import (
	"strings"
)

func Check(variants *[]string, str string) []int {
	variant := *variants
	capacity := len(str) / len(variant)
	if capacity == 0 {
		return []int{}
	}
	arrIndex := make([]int, 0, capacity)
	index := 0
	equal := false
	strArr := strings.Split(str, "")
	j := 0
	for i := 0; i < len(str); i++ {
		if strArr[i] == variant[j] {
			if j == len(variant)-1 && equal {
				arrIndex[index] = i - len(variant) - 1
				index++
			} else {
				j++
				equal = true
			}
		} else {
			equal = false
			j = 0
		}
	}
	return arrIndex
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
