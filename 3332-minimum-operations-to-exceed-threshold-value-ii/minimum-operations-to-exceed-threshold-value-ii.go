type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func minOperations(a []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, val := range a {
		heap.Push(h, val)
	}
    fmt.Println(h)
	for i := 0; ; i++ {
		x := heap.Pop(h).(int)
		if x >= k {
			return i
		}
		y := heap.Pop(h).(int)
		t := int64(2 * x + y) // test overflow
		if t < int64(k) {
			heap.Push(h, int(t))
		} else {
			heap.Push(h, k)
		}
	}
	return -1
}