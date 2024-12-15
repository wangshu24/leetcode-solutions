// helper function to get min and max from a map
func getMinMax(m map[int]int) (min, max int) {
    if len(m) == 0 {
        return 0, 0
    }
    min = math.MaxInt32
    max = math.MinInt32
    
    for val := range m {
        if val < min {
            min = val
        }
        if val > max {
            max = val
        }
    }
    return
}


func continuousSubarrays(nums []int) int64 {
    count := int64(0)
    freq := make(map[int]int)
    
    start := 0
    for end, n := range nums {
        freq[n]++
        
        minVal, maxVal := getMinMax(freq)
        for maxVal - minVal > 2 {
            if freq[nums[start]]--; freq[nums[start]] == 0 {
                delete(freq, nums[start])
            }
            minVal, maxVal = getMinMax(freq)
            start++
        }
        
        count += int64(end - start + 1)
    }
    
    return count
}