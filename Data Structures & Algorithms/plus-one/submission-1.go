func plusOne(digits []int) []int {
   for i := len(digits) - 1; i >=0; i-- {
	if digits[i] < 9 {
		digits[i]++
		return digits // これってforの途中でもreturnだったらforループは終了するっけ？
	}
	digits[i] = 0	
   } 

   return append([]int{1}, digits...)
}
