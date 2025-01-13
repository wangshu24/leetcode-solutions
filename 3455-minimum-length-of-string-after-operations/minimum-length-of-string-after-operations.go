func minimumLength(s string) int {
    chars := make([]int, 26)
    for _, char := range s {
        chars[char-97]++
    } 
    fmt.Println(chars)
    count := 0
    for _, val := range chars {
      
        if val == 0 {continue} 
        if val < 3 {
            count+=val
        } else if val %2 == 0 {
            count+=2
        } else  {
            count++
        }
        fmt.Println(count)
    }

    return count
}