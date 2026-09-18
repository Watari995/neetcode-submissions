import "slices"

func minEatingSpeed(piles []int, h int) int {
	// left, right で binary searchをする // rightはpilesのmax、これが最速のスピードだからこれ以上速度を上げることはできない
	left := 1 
	right := slices.Max(piles)
	result := right

	// kを一つ決めてその中でチェックしていく
	for left <= right {
		k := (left + right) / 2
		if canEat(piles, h, k) {
			result = k
			right = k - 1 // さらに小さいkがある可能性があるので減らす
		} else {
		// 食べきれないケースは、一つずつ増やしていく
		left = k + 1
		}
	}
	return result
}

func canEat(piles []int, h, k int) bool {
	hours := 0 // これに足していく
	for _, pile := range piles {
		hours += (pile + k - 1) / k
	}

	return h >= hours
}
