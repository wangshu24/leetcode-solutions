func findThePrefixCommonArray(A []int, B []int) []int {
    freq  := make([]int, len(A)+1)
    res := make([]int, len(A))
    comm := 0
    fmt.Println(res)
    for i:=0; i < len(A); i++ {
        freq[A[i]]++
        if freq[A[i]] == 2 {comm++}
        freq[B[i]]++
        if freq[B[i]] == 2 {comm++}
        res[i] = comm
    }

    return res
}