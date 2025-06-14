package twosum

func TwoSum(nums []int, target int) []int {
	count := make(map[int]int)

	for i, num := range nums {
		complememt := target - num

		if v, ok := count[complememt]; ok {
			return []int{v, i}
		}
		count[num] = i
	}
	return []int{-1, -1}
}
