func twoSum(nums []int, target int) []int {
// key = value, value = index for nums
   indexMap := make(map[int]int) 
   for i, n := range nums {
    indexMap[n] = i
   }

   for i, v := range nums {
    compl := target - v

    if value, ok := indexMap[compl]; ok {
        if value != i {
        return []int{i, value}
        }
    }
   }
   return []int{}
}
