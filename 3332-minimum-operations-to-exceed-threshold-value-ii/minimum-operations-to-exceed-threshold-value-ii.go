
func minOperations(nums []int, k int) int {
    h := &MinIntHeap{}
    heap.Init(h)
    count := 0
    for _, num := range nums {
        if num < k {
            heap.Push(h, num)
        }
    }
    
    for h.Len() > 0 {
        a := heap.Pop(h).(int)
        if h.Len() == 0 {
            count++
            break
        }
        b := heap.Pop(h).(int)
        num := a*2+b
        if num < k {
            heap.Push(h, num)
        }
        count++
    }
    return count
}





type MinIntHeap []int
func(h MinIntHeap) Len() int { return len(h) }
func(h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func(h MinIntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func(h *MinIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func(h *MinIntHeap) Pop() interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return x
}