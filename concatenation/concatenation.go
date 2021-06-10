package concatenation

import (
	"strings"
)

func Check(variant string, str string) []int {
	arrVariant := strings.Split(variant, "")
	var arrIndex []int
	index := 0
	equal := false
	strArr := strings.Split(str, "")
	j := 0
	for i := 0; i < len(str); i++ {
		if strArr[i] == arrVariant[j] {
			if j == len(arrVariant)-1 && equal {
				arrIndex = append(arrIndex, i-(len(arrVariant)-1))
				index++
				i++
			} else {
				j++
				equal = true
			}
		} else {
			if equal {
				i--
				if i < 0 {
					i = 0
				}
			}
			equal = false
			j = 0
		}
	}
	if arrIndex == nil {
		return []int{}
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
