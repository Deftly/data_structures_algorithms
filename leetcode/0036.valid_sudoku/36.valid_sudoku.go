package data_structures_algorithms

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]bool
	var squares [3][3][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]

			if cell != '.' {
				value := int(cell-'0') - 1

				if rows[i][value] || columns[j][value] || squares[i/3][j/3][value] {
					return false
				}

				rows[i][value] = true
				columns[j][value] = true
				squares[i/3][j/3][value] = true
			}
		}
	}
	return true
}
