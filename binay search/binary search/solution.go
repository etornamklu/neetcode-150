package binary_search

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for r >= l {
		m := l + (r-l)/2

		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		} else {
			return m
		}

	}

	return -1

}
