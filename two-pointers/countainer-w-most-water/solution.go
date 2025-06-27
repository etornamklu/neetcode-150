package countainer_w_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for r > l {
		area := min(height[l], height[r]) * (r - l)

		if area > res {
			res = area
		}
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}

	return res
}
