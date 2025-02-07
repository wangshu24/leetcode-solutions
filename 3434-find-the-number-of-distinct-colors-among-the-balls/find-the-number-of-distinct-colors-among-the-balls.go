func queryResults(limit int, queries [][]int) []int {
    colMap := map[int]int{}
    ballMap := map[int]int{}
    res := make([]int, len(queries))

    for i:=0; i < len(queries); i++{
        col := queries[i][1]
        ball := queries[i][0]

        // if prevCol, ok := ballMap[ball]; ok {
        //     colMap[prevCol]--
        //     if colMap[prevCol] == 0 {delete(colMap, prevCol)}
        // }
        if ballMap[ball] > 0 {
            colMap[ballMap[ball]]--
            if colMap[ballMap[ball]] == 0 {
                delete(colMap, ballMap[ball])
            }
        } 

        ballMap[ball] = col 
        colMap[col]++
        res[i] = len(colMap)
    }

    fmt.Println(colMap)
    return res
}

