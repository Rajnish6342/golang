package setmatrixzeros

func setZeroes(matrix [][]int) {
	rowSize := len(matrix)
	colSize := len(matrix[0])
	row := make([]bool, rowSize)
	col := make([]bool, colSize)

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}

		}
	}
	for i := 0; i < rowSize; i++ {
		if row[i] {
			for j := 0; j < colSize; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < colSize; i++ {
		if col[i] {
			for j := 0; j < rowSize; j++ {
				matrix[j][i] = 0
			}
		}
	}

}
