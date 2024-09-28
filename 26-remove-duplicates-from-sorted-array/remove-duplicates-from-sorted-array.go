func removeDuplicates(nums []int) int {
    dup := 0
    unique := 1
    for i:=0; i< len(nums); i++ {
        if i == 0 {
            dup = nums[0]
        } else {
            if nums[i] != dup{
                nums[unique] = nums[i]
                dup = nums[i]
                unique++
            } 
        }
    }
    return unique
}