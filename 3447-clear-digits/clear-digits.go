func clearDigits(s string) string {
    var builder strings.Builder
    mark := make([]int, len(s))
    stack := make([]int, 0)
    for i:=0; i < len(s); i++ {
        if s[i] >= 48 && s[i] <= 57 {
            mark[i] = 1
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack,i)
        }
    }
    
    for i:=0; i < len(stack); i++ {
        builder.WriteByte(s[stack[i]])
    }
    

    return builder.String()
} 