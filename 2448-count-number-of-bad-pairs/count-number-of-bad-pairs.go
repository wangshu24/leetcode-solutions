func countBadPairs(nums []int) int64 {
    var totalPairs int = (len(nums) * (len(nums) - 1))/2
    var goodCount int
    var goods = map[int]int{}
    for i:=0; i<len(nums); i++ {
        val := i - nums[i]
        goodCount += goods[val]
        goods[val]++
    }
    
    return int64(totalPairs - goodCount)
}