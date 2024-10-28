func longestSquareStreak(nums []int) int {
    numMap := map[int]int{}
    slices.Sort(nums)
    str:=-1

    for _,num := range nums {
        sqrt := int(math.Sqrt(float64(num)))
        if sqrt*sqrt == num {
            if val, exist := numMap[sqrt]; exist {
                numMap[num] = val + 1
                if numMap[num] > str {
                    str = numMap[num]
                }
            } else {
                numMap[num] = 1
            }
        } else {
            numMap[num] = 1
        }
    }

    return str
}