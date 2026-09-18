type Pair struct {
	value string
	timestamp int
}

type TimeMap struct {
	store map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], Pair{
		value: value,
		timestamp: timestamp,
	}) 
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// timestamp以下の値で最新のやつ(一番右)を取り出す 
	pairs := this.store[key]
	left, right := 0, len(pairs) - 1
	result := ""

	for left <= right {
		mid := (left + right) / 2

		if pairs[mid].timestamp <= timestamp {
			result = pairs[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
