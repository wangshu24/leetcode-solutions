type MinHeap []int

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(a,b int) bool {return h[a] < h[b]}
func (h MinHeap) Swap(a,b int) {h[a], h[b] = h[b], h[a]}
func (h *MinHeap) Push(i interface{}) {
    *h = append(*h, i.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := (*h)[len(*h)-1]
    *h =  (*h)[:len(*h)-1]
    return old
}

func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}
    for _, num := range nums {
        if h.Len() >= k && (*h)[0] < num {
            heap.Pop(h)
            heap.Push(h, num)
        } else if h.Len() < k {
            heap.Push(h, num)
        } 
    }

    fmt.Println(*h)
    return (*h)[0]
}