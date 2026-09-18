func hasDuplicate(nums []int) bool {
   numMap := make(map[int]bool)
   for _, v := range nums {
	if _, ok := numMap[v]; ok {
		return true
	}
	numMap[v] = true
   }
   return false
}
