func canConstruct(s string, k int) bool {
    if len(s) < k {return false}

    //solution with hashmap
    chars := make([]int, 26)
    for _, char := range s {
        chars[char-97]++
    }
    
    odds := 0
    for _,val := range chars {
        if val % 2 !=  0 {odds++}
        if odds > k {return false}
    }

    return true
}