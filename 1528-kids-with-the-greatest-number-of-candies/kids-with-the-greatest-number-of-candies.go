func kidsWithCandies(candies []int, extra int) []bool {
    res:= make([]bool, len(candies))
    max := slices.Max(candies)
    
    for i:=0; i < len(candies);i++{
        if candies[i]+extra >=max {
            res[i] = true
        } else {res[i] = false}
    }

    return res
}