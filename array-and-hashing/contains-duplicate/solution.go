package containsduplicate

func ContainsDuplicate(nums []int) bool {
	tracker := make(map[int]bool) // add some notes on how to create maps with make

	for _, num := range nums { // value, exits
		if _, ok := tracker[num]; ok {
			return true
		}
		tracker[num] = true
	}

	return false
}
