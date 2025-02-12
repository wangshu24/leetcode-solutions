func maximumSum(nums []int) int {
    total := map[int][]int{}
    slices.Sort(nums)
    for _, num := range nums {
        sum := sumDigits(num)
        total[sum] = append(total[sum], num)
    }

    res:=-1
    for _, list := range total {
        if len(list) >1 {
            res = max(res, list[len(list)-1] + list[len(list)-2] )
        }
    }
 
    return res
}

func sumDigits(num int) int {
    res := 0
    for num > 0 {
        res+=num%10
        num = num/10
    }
    return res
}

func max(a, b int) int {
    if a > b {return a}
    return b
}