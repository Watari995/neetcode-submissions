import "slices"

func threeSum(nums []int) [][]int {
	// ソートする
	slices.Sort(nums)

	res := [][]int{}
	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := n + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{n, nums[l], nums[r]})
				// ここで足していかないとずっと同じ値でループしてしまうので
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			if sum > 0 {
				r--
			} 
			if sum < 0{
				l++
			}
		}
	}
	return res
}
