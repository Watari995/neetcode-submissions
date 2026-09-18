func searchMatrix(matrix [][]int, target int) bool {
	// i番目とi+1番目の範囲を見る、targetがその範囲内であればi番目の行をhalfのbinary searchの方法で見る
	rowLen := len(matrix)

	for i := range rowLen - 1 {
		if matrix[i][0] == target || matrix[i+1][0] == target {
			return true
		}
		if matrix[i+1][0] > target {
			// matrix[i] にtargetがあるということ
			left := 0
			right := len(matrix[i]) - 1
			for left <= right {
				half := (left+right) / 2
				if matrix[i][half] == target {
					return true
				}
				if matrix[i][half] > target {
					right = half - 1
				} else {
					left = half + 1
				}
			}
		} else {
			continue // ないので次のi+1をチェックする
		}
	}
	last := rowLen - 1
left := 0
right := len(matrix[last]) - 1
for left <= right {
    half := (left + right) / 2
    if matrix[last][half] == target {
        return true
    }
    if matrix[last][half] > target {
        right = half - 1
    } else {
        left = half + 1
    }
}
return false
	return false
}
