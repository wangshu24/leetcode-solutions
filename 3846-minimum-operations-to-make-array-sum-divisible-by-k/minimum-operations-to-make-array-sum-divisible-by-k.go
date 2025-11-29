func minOperations(nums []int, k int) int {
    var sum int 
    for _,i := range nums {
        sum+=i
    }
    return sum%k
}