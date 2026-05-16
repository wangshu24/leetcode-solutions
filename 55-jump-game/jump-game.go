func canJump(nums []int) bool {
    dest := len(nums)-1
    for i:= len(nums)-2; i >= 0; i-- {
        if nums[i] + i >= dest {
            dest = i
        } 
    }
    
    return dest == 0
}