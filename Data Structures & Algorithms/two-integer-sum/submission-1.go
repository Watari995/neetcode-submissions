func twoSum(nums []int, target int) []int {
    // numsの一番左から左+1, 左+2, 左+3... len(nums) - 1 まで検証していく
	// だんだん検証の範囲を左にずらして小さくする
	for i := 0;i < len(nums) - 1; i++ {
		for k := i + 1;k < len(nums); k++ {
			if nums[i] + nums[k] == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}
