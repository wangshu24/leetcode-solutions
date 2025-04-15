func maxCompatibilitySum(students [][]int, mentors [][]int) int {
    var (
        res int
        ROW = len(students)
        paired  = map[int]bool{}
    )

    var dfs func(i, sum int, path map[int]bool)

    dfs = func(i, sum int, path map[int]bool) {
        if i == ROW {
            if res == 0 || sum > res {
                res = sum
            }
            return
        }

        for j:=0; j < ROW; j++ {
            if !path[j] {
                path[j] = true
                dfs(i+1, sum+comp(students[i], mentors[j]), path)
                path[j] = false
            }
        }
    }
    
    dfs(0,0, paired)
    return res
}

func comp(stu, men []int) int {
    var res int
    for i:=0; i < len(stu); i++ {
        if stu[i] == men[i] {
            res++
        }
    }
    return res
}
