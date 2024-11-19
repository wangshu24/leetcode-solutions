func maximumSubarraySum(nums []int, k int) int64 {
    res, being, tempMax := 0, 0, 0
    kMap:= map[int]int{}

    for i:=0; i < len(nums); i++ {
       if kMap[nums[i]] == 0 {
            kMap[nums[i]]++
            tempMax+= nums[i]

            if i - being + 1==k {
                res = max(res, tempMax)
                tempMax -= nums[being]
                kMap[nums[being]]--
                being++
            }        
       } else {
            for being < i && nums[being] != nums[i]{
                tempMax -= nums[being]
                kMap[nums[being]]--
                being++
            }
            being++
       }
    }

    return int64(res)
}

func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}