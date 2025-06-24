package twosumiiinputarrayissorted

func TwoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1

	for r > l {
		if numbers[r]+numbers[l] > target {
			r--
		} else if numbers[r]+numbers[l] < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{-1, -1}

}
