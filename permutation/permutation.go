package permutation

func Permutate(input *[]string, l int, h int, answers *[]string) {
	array := *input
	if l == h-1 {
		*answers = append(*answers, sum(&array))
	} else {
		for i := l; i < h; i++ {
			array[l], array[i] = array[i], array[l]
			Permutate(&array, l+1, h, answers)
			array[l], array[i] = array[i], array[l]
		}
	}
}

func sum(input *[]string) string {
	array := *input
	result := ""
	for _, v := range array {
		result += v
	}
	return result
}
