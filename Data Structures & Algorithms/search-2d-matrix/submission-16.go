func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	left, right := 0, rows*cols-1 // これはindexが入っている

	for left <= right {
		half := (left+right) / 2 
		row := half / cols
		col := half % cols

		value := matrix[row][col]
		if value == target {
			return true
		}
		if value > target {
			right = half - 1 
		} else {
			left = half + 1
		}
	}
	return false
}
