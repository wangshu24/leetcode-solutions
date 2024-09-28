func removeElement(nums []int, val int) int {
    modified := slices.DeleteFunc(nums, func (n int) bool {
        return n == val
    })
    
    return len(modified)
}