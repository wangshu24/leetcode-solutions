type KthLargest struct {
    k int
    heap Heap
}

type Heap []int 

func (h Heap) Len() int {return len(h)}
func (h Heap) Less(a,b int) bool {return h[a] < h[b]}
func (h Heap) Swap(a,b int) {h[a], h[b] = h[b], h[a]}
func (h *Heap) Push(a interface{}) {
    (*h) = append((*h), a.(int))
}

func(h *Heap) Pop() interface{} {
    old := (*h)[len(*h)-1]
    (*h) = (*h)[:len(*h)-1]
    return old
}

func Constructor(k int, nums []int) KthLargest {
    minHeap := &Heap{}
    heap.Init(minHeap)
    klargest := KthLargest{k : k, heap: *minHeap} 
    for _, num := range nums {
        klargest.Add(num)
    }
    return klargest
}


func (this *KthLargest) Add(val int) int {
    if this.heap.Len() < this.k {
        heap.Push(&this.heap, val)
    } else if val > this.heap[0] {
        heap.Pop(&this.heap)
        heap.Push(&this.heap, val)
    }

    return this.heap[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */