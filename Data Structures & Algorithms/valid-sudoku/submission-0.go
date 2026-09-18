func isValidSudoku(board [][]byte) bool {
	// その3x3のマス全部と、縦と横全て 1-9揃っている必要がある
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// 初期化をする
	for i :=0; i< 9; i++{
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// それで順番にチェックしていく
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			v := board[i][j]
			if v == '.' { // byteなのでシングルクォート
				continue // .の時は処理をskipする
			}
			if rows[i][v] || cols[j][v] || boxes[(i/3)*3+j/3][v] {
				return false
			}
			rows[i][v] = true
			cols[j][v] = true
			boxes[(i/3)*3+j/3][v] = true
		}
	}
	return true
}
