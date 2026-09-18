func isValid(s string) bool {
   pairs := map[rune]rune{
	')' : '(',
	'}' : '{',
	']' : '[',
   } 

   stack := []rune{}

   for _, v := range s {
	if value, ok := pairs[v]; ok {
	if len(stack) == 0 {
		return false
	}
		top := stack[len(stack)-1]
		if value != top {
			return false
		}
		stack = stack[:len(stack)-1]
	} else {
		stack = append(stack, v)
	}
   }
   return len(stack) == 0
}
