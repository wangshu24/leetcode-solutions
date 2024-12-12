import ("slices")

func pickGifts(gifts []int, k int) int64 {
    res := 0
    slices.Sort(gifts)
    for i:=0; i < k; i++ {
        gifts[len(gifts)-1] = int(math.Sqrt(float64(gifts[len(gifts)-1])))
        slices.Sort(gifts)
    }

    for _,val := range gifts {
        res+=val
    }

    return int64(res)
}