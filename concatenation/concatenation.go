package concatenation

import (
	"strings"
)

func Check(variant string, str string) []int {
	arrVariant := strings.Split(variant, "")
	var arrIndex []int
	equal := false
	equalIndex := -1
	strArr := strings.Split(str, "")
	j := 0
	for i := 0; i < len(str); i++ {
		if strArr[i] == arrVariant[j] {
			if j == len(arrVariant)-1 && equal {
				arrIndex = append(arrIndex, i-(len(arrVariant)-1))
				j = 0
				i++
			} else {
				if equalIndex == -1 {
					equalIndex = i + 1
				}
				j++
				equal = true
			}
		} else {
			if equal {
				i = equalIndex
				equalIndex = -1
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
