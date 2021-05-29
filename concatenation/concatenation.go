package concatenation

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckStr(variants *[]string, str string) []int {
	arr := *variants
	arrStr := strings.Split(str, "")
	var stack Stack

	for i := 0; i < len(arr)-1; i++ {
		j := 0
		arrVal := strings.Split(arr[i], "")
		k := 0

		if arrVal[k] != arrStr[j] {
			stack.Pop()
			k = 0
			j++
			if j > len(arrStr)-1 {
				break
			}
		} else {
			if k == 0 {
				stack.Push(strconv.Itoa(j))
				j++
				if j > len(arrStr)-1 {
					if k != len(arrVal[k])-1 {
						stack.Pop()
					}
					break
				}
				k++
			}
		}
	}

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
	// for _, variant := range *variants {
	// 	for _, v := range strings.Split(variant, "") {
	// 		if arr[index] == v {
	// 			results[resIndex] = index
	// 			index = index + 1
	// 			if index == len(str) {
	// 				return results
	// 			}
	// 		} else {
	// 			index = 0
	// 		}
	// 	}
	// }
	return []int{2}
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
