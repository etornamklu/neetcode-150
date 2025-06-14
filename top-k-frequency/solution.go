package topkfrequency

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num]++
	}

	for num, c := range count {
		freq[c] = append(freq[c], num)
	}

	res := make([]int, 0, k)
	for i := len(freq) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, freq[i]...)
	}

	return res[:k]
}
