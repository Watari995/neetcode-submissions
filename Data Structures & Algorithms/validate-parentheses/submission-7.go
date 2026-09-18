func isValid(s string) bool {
	pairs := map[byte]byte {
		')' : '(',
		']' : '[',
		'}' : '{',
	}
	stack := []byte{} // 空スライスで初期化

	for i:=0; i<len(s); i++ {
		v := s[i] // 値を代入
		if value, ok := pairs[v]; ok {
			// 閉じ括弧があった場合
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1] // 一番最近積んだもの
			stack = stack[:len(stack)-1]
			if top != value {
				return false
			}
		} else { // 閉じ括弧ではない(開き括弧)のケース
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
