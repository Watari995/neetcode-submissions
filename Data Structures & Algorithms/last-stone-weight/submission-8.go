
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i,j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x any) {
	*h = append(*h,x.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h 
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)
		// 値が一緒だったら何もしないけど、違ったら差分を積む。
		if y != x {
			heap.Push(&h, y-x)
		}
	}
		if h.Len() == 0 {
		return 0
		}
	return h[0]
}
