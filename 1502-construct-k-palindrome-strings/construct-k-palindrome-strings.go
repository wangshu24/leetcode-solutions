func canConstruct(s string, k int) bool {
    if len(s) < k {return false}

    chars := map[string]int{}
    for _, char := range s {
        chars[string(char)]++
    }
    
    fmt.Println(chars)
    odds := 0
    for _,val := range chars {
        if val % 2 !=  0 {odds++}
        if odds > k {return false}
    }

    return true
}