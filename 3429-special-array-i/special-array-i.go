func isArraySpecial(nums []int) bool {
    if len(nums) == 1 {return true}
    
    tmp := nums[0] % 2
    fmt.Println(tmp)

    for i:= 1; i < len(nums); i++ {
        if nums[i]%2 == tmp {
            return false
        }
        tmp = nums[i]%2
    } 

    return true
}