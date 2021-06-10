package concatenation

import (
	"strings"
)

func Check(variant string, str string) []int {
	arrVariant := strings.Split(variant, "")
	var arrIndex []int
	equal := false
	strArr := strings.Split(str, "")
	j := 0
	for i := 0; i < len(str); i++ {
		if strArr[i] == arrVariant[j] {
			if j == len(arrVariant)-1 && equal {
				arrIndex = append(arrIndex, i-(len(arrVariant)-1))
				j = 0
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
