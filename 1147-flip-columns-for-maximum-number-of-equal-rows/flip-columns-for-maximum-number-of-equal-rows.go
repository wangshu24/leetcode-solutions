func maxEqualRowsAfterFlips(mat [][]int) int {
    bitMap := map[string]int{}

    for _,row := range mat {
        pattern := make([]byte, len(row))
        if row[0] == 0 {
            for i,bit := range row{
                pattern[i] = byte(bit) + '0'
            }
            fmt.Println(pattern)
        } else {
            for i,bit := range row{
                pattern[i] = byte(bit)^1 + '0'
            }
            fmt.Println(pattern)
        }
        
        bitMap[string(pattern)]++
    }
    fmt.Println(bitMap)
    max := 0
    for _,val := range bitMap {
       if val > max {
        max = val
       }
    }
    return max
}

// func max(a,b int ) int {
//     if a > b {
//         return a
//     } 
//     return b
// }