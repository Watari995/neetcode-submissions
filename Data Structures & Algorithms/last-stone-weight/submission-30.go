type HeapInt []int

func (h HeapInt) Len() int { return len(h) }
func (h HeapInt) Less(i, j int) bool { return h[i] > h[j] }
func (h HeapInt) Swap (i, j int) { h[i], h[j] = h[j], h[i]}
func (h *HeapInt) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *HeapInt) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := HeapInt(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if x != y {
			heap.Push(&h, x-y)	
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}
