func longestConsecutive(nums []int) int {
	// setを使用して、1発で当てる
	// for文を2回回すことで、O(n)にすることができる
	set := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := set[n]; !ok {
			set[n] = struct{}{}
		}
	}

	// 一番小さい値から順番にチェックしていく
	maxLength := 0
	for _, n := range nums {
		// n-1をしていって、それがあればさらに-1をして最小値を探す、なければそれが最小値なので、そこからsetの探索スタート
		if _, ok := set[n-1]; !ok {
			length:= 1
			for {
				if _, ok := set[n+1]; ok {
					length++
					n++
				} else {
					break
				}
			}
			maxLength = max(maxLength, length)
		}
	}
	return maxLength
}
