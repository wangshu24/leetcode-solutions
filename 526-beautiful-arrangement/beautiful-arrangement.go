func countArrangement(n int) int {
    var (
        res int
        backtrack func(i int)
        visit = make(map[int]bool)
    )    

    backtrack = func(i int) {
        if i == n {
            res++
            return
        }

        for j:=0; j < n; j++ {
            if !visit[j+1] && beautiful(i+1, j+1) {
                fmt.Println(i, j+1)
                visit[j+1] = true
                backtrack(i+1)
                visit[j+1] = false
            }
        } 
    }

    backtrack(0)
    return res
}

func beautiful(a, b int) bool {
    return a%b == 0 || b%a == 0
}