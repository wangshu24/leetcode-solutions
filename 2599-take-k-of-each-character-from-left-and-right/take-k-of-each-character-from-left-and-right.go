func takeCharacters(s string, k int) int {
    remove, l := 0, 0
    kmap := make([]int, 3)
 
    if len(s) < k*3 {return -1}
    for i:=0; i < len(s); i++{
        kmap[s[i]-97]++
    }
    
    if kmap[0] < k || kmap[1] < k || kmap[2] < k {return -1}

    for r:=0; r< len(s); r++{
        kmap[s[r]-97]--

        if slices.Min(kmap) < k {
            for slices.Min(kmap) < k {
                kmap[s[l]-97]++
                l++
            }      
        } else {
            fmt.Println(r-l)
            remove = max(remove, r-l+1)
        }
    }

    return len(s)-remove
}

func max(a,b int) int{
    if a > b {return a}
    return b
}
