func twoSum(nums []int, target int) []int {
   numMap := make(map[int]int)
   for i, n := range nums {
    numMap[n] = i
   } 

   for i, v := range nums {
    complement := target - v
    value, ok := numMap[complement]
    
    if value != i  && ok {
        return []int{i, value}
    }
   }
   return []int{}
}
