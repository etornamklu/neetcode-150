package validsudoku

func IsValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	squares := make([]int, 9)

	for r := range 9 {
		for c := range 9 {

			if board[r][c] == '.' {
				continue
			}

			val := board[r][c] - '1'
			bit := 1 << val
			squareIdx := (r/3)*3 + c/3

			if rows[r]&bit != 0 || cols[c]&bit != 0 ||
				squares[squareIdx]&bit != 0 {
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			squares[squareIdx] |= bit
		}
	}

	return true
}

func IsValidSudokuBruteForce(board [][]byte) bool {

	// validating rows
	for r := range board {
		counter := make(map[byte]bool)
		for c := range board[r] {

			v := board[r][c]

			if v == '.' {
				continue
			}

			if _, ok := counter[v]; ok {
				return false
			}

			counter[v] = true
		}
	}

	//validating columns
	for c := range board {
		counter := make(map[byte]bool)
		for r := range board[c] {

			v := board[r][c]

			if v == '.' {
				continue
			}

			if _, ok := counter[v]; ok {
				return false
			}

			counter[v] = true
		}
	}

	//validating 3x3 squares
	for square := range 9 {
		counter := make(map[byte]bool)
		for i := range 3 {
			for j := range 3 {

				r := (square/3)*3 + i
				c := (square%3)*3 + j

				v := board[r][c]

				if v == '.' {
					continue
				}

				if _, ok := counter[v]; ok {
					return false
				}

				counter[v] = true
			}
		}

	}

	return true

}
