func minimumOperations(nums []int) int {
    var count int
    for _,i := range nums {
        if i%3 != 0 || i < 3  {
            count++
        }
    }

    return count
}