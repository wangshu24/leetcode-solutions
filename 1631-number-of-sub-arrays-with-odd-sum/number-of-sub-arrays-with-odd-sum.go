const mod = 1e9 + 7 
func numOfSubarrays(arr []int) int {
    var res int 
    prefixSum, oddCount, evenCount := 0, 0, 1

    for _, num := range arr {
        prefixSum += num
        if prefixSum%2 == 0 {
            res = (res + oddCount) % mod
            evenCount++
        } else {
            res = (res + evenCount) % mod
            oddCount++
        }
    }

    return res
}