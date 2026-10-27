import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	for i := range nums {
		// 先頭ではないのでskipとする
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			if sum < 0 {
				l++
			}
			if sum > 0 {
				r--
			}
		}
	}
	return res
}
