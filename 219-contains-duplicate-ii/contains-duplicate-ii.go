func containsNearbyDuplicate(nums []int, k int) bool {
    col := map[int][]int{}
    result := false
    for i:=0; i < len(nums);i++ {
        col[(nums[i])] = append(col[(nums[i])], i)
    }
    
    for _, v := range col {
        if len(v) >1 {
            for i,j := range v {
                if i+1<len(v) {
                    fmt.Println(v[i+1])
                    fmt.Println(j)
                    if ((v[i+1] - v[i]) <= k) {
                        result = true
                        break
                    }
                }
            }
        }
    }

    return result
}