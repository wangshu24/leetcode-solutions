func combine(n int, k int) [][]int {
    var res [][]int

    var backtrack func(i int, set []int)
    backtrack = func(i int, set []int) {
        if len(set) == k {
            dst := make([]int, k)  
            copy(dst,set)
            res = append(res, dst)
            return
        } else if i > n || len(set) > k {
            return
        }

        set = append(set, i)
        backtrack(i+1, set)
        set = set[:len(set)-1]
        backtrack(i+1, set)
    }

    backtrack(1, make([]int,0))
    return res
}