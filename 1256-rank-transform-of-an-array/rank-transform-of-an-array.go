func arrayRankTransform(arr []int) []int {
    orig := slices.Clone(arr)
 
    slices.Sort(arr)
    fmt.Println(arr)
    rankMap := map[int]int{}
    rank := 1
    for i:=0; i < len(arr); i++ {
        if rankMap[arr[i]] ==  0 {
            rankMap[arr[i]] = rank
            rank++
            }
        
    }
    fmt.Println(rankMap)
    for i:=0; i < len(arr); i++ {
        orig[i] = rankMap[orig[i]]
    }
    return orig
}