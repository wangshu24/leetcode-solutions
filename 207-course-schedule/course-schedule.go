func canFinish(numCourses int, prerequisites [][]int) bool {
    degree := make([]int, numCourses)
    adj := make(map[int][]int)

    for _, val := range prerequisites {
        degree[val[0]]++
        adj[val[1]] = append(adj[val[1]], val[0])
    }
    
    var cycle func() bool
    cycle = func() bool {
        q := []int{}
        count := 0
        for ind, val := range degree {
            if val == 0 {
                q = append(q, ind)
                count++
            }
        }

        for len(q) > 0 {
            x := q[0] 
            q = q[1:]
            for _, val := range adj[x] {
                degree[val]--
                if degree[val] == 0 {
                    count++
                    q = append(q, val)
                }
            }
        }
        if count == numCourses {
            return true
        }
        return false
    }

    return cycle()
}