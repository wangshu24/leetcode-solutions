func smallestNumber(pattern string) string {
    stack, res := []byte{}, []byte{}
    
    for i:=0; i < len(pattern)+1; i++ {
        stack = append(stack, byte(i+'1'))

        if i == len(pattern) || pattern[i] == 'I' {
            for j := len(stack); j > 0; j-- {
                res = append(res, stack[j-1])
            }
            stack  = stack[:0]
        }
    }
    return string(res)
}