func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
        // stackから2つ取り出す
        b := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        a := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 計算してstackに戻す
		switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
		}
    	} else {
        // 数字の場合はintに変換してstackに積む
        num, _ := strconv.Atoi(token)
        stack = append(stack, num)
    	}
	}
	return stack[0] // 先頭を返す (一つしか残ってないはずなので)
}
