const mod = 1e9+7
func numOfSubarrays(arr []int) int {
    var res int
    prefSum, odd, even := 0, 0, 1

    for _,num := range arr {
        prefSum += num 
        if prefSum %2 == 0 {
            res += odd
            even++
        } else {
            res+= even
            odd++
        }
    }
    return res%mod
}