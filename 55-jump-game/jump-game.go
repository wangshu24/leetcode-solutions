func canJump(nums []int) bool {
    jumps := make([]int, len(nums))
    for ind, val := range nums {
        jumps[ind] = val
    }
    end := make([]int, len(nums))
    
    var back func( int) bool 
    back = func(cur int) bool {
        if end[cur] == -1 {
            return false
        }
        if cur >= len(nums)-1 || end[cur] == 1 {
            return true
        }
        
        can := false
        for i:=1; i <= jumps[cur]; i++ {
            can = can || back(cur+i)
        }
        if !can {
            end[cur] = -1
        }        
        return can
    }

    return back(0)
}