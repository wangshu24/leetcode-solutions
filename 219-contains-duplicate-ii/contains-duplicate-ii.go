func containsNearbyDuplicate(nums []int, k int) bool {
    col := map[int][]int{}
    result := false
    for i:=0; i < len(nums);i++ {
        col[(nums[i])] = append(col[(nums[i])], i)
    }
    
    for _, v := range col {
        if len(v) >1 {
            for i := 0; i < len(v)-1; i++ {
                fmt.Println(v[i+1])
                fmt.Println(v[i])
                if (int(v[i+1]) - int(v[i])) <= k {
                    result = true
                    break
                }
            } 
        }
    }

    return result
}