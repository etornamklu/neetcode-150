package search_matrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	r, c := 0, n-1

	for r < m && c >= 0 {
		// If the current element is greater than the target, move left (stay in the same row).
		if matrix[r][c] > target {
			c--
			// If the current element is less than the target, move down (stay in the same column).
		} else if matrix[r][c] < target {
			r++
		} else {
			return true
		}
	}
	return false
}
