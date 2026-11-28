import "slices"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, slices.Max(piles)
	for l < r {
		mid := l + (r-l) / 2 // prevent overflow 
		hours := 0
		for _, p := range piles {
			// add hour and check if all added number is within k
			k := p / mid
			if p % mid != 0 {
				k++ // ceilup
			}
			hours += k
		}
		if hours <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
