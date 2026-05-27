type Heap []Word

type Word struct {
    word string
    count int
}

func (h Heap) Len() int {return len(h)}
func (h Heap) Less(a, b int) bool {
    if h[a].count == h[b].count {
		return h[a].word > h[b].word
	}
	return h[a].count < h[b].count
}
func (h Heap) Swap(a, b int) {h[a], h[b] = h[b], h[a]}
func (h *Heap) Push(i interface{}) {
    *h = append(*h, i.(Word))
}
func (h *Heap) Pop() interface{} {
    old  := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return old
}

func topKFrequent(words []string, k int) []string {
    h := &Heap{}
    heap.Init(h)
    wordmap := make(map[string]int)
    for _, word := range words {
        wordmap[word]++
    }

    for key, val := range wordmap {
        heap.Push(h, Word{word: key, count: val})
        if h.Len() > k {
            heap.Pop(h)
        }
    } 
    res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(h).(Word)
		res[i] = item.word
	}

    return res
}