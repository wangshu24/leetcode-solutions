func totalFruit(fruits []int) int {
    start, res  := 0, 0 
    ft := map[int]int{}

    for i:=0; i < len(fruits); i++ {
        ft[fruits[i]]++
        
        for len(ft) > 2 {
            ft[fruits[start]]--
            if ft[fruits[start]] == 0 {
                delete(ft, fruits[start])
            }
            start++
        }
        res = max(res, i - start+1)
    }

    return res
}

func max(a, b int) int {
    if a > b {return a}
    return b
}