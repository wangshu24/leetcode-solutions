func canPartition(nums []int) bool {
    if len(nums) == 1 {return false}
    var sum int

    for _,val := range nums {
        sum+=val
    } 
    if sum%2 != 0 {return false}
    
    dp := make([]bool, (sum/2)+1)
    dp[0] = true
    for _, i := range nums {
        for j:=sum/2; j >= i; j--{
            if dp[j-i] {
            dp[j] = true
            }
        }
    }

    return dp[sum/2]
}