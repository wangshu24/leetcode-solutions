// func minGroups(intervals [][]int) int {
//     start := intervals[0][0]
//     end := intervals[0][1]
//     endArr := []int{}
//     for i:=0; i <len(intervals);i++{
//         if intervals[i][0] < start { start = intervals[i][0] }
//         if intervals[i][1] > end { end = intervals[i][1] } 
//     }
//     minGroup := len(intervals)
//     for i:=start; i<end;i++{
        
//     }
    
//     fmt.Println(start)
//     fmt.Println(end)
//     return 0
// }


import (
    "container/heap"
    "sort"
)
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
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
func minGroups(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    pq := &MinHeap{}
    heap.Init(pq)
    for _, interval := range intervals {
        start, end := interval[0], interval[1]
        if pq.Len() > 0 && (*pq)[0] < start {
            heap.Pop(pq)
        }
        heap.Push(pq, end)
    }
    return pq.Len()
}