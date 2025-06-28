package koko_eating_bananas

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0

	for _, p := range piles {
		if p > r {
			r = p
		}
	}

	for r > l {
		m := l + (r-l)/2
		totalHours := 0

		for _, pile := range piles {
			totalHours += (pile + m - 1) / m // Equivalent to ceil(pile / m)
		}

		if totalHours > h {
			l = m + 1 // Need to eat faster
		} else {
			r = m // Can try slower speeds
		}
	}
	return l // The minimum speed that allows Koko to eat all bananas in h hours

}
