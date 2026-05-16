func canFinish(numCourses int, prerequisites [][]int) bool {
    indegree := make([]int, numCourses)
    adj := make(map[int][]int)
    for _, preq := range prerequisites {
        a, b := preq[0], preq[1]
        adj[b] = append(adj[b], a)
        indegree[a]++
    }
    fmt.Println(adj, indegree)
    
    var cycle func() bool
    cycle = func() bool {
        q := []int{}
        count := 0
        for ind, val := range indegree {
            if val == 0 {
                q = append(q, ind)
                count++
            }
        }

        for len(q) > 0 {
            c := q[0]
            q = q[1:]
            for _, v := range adj[c] {
                indegree[v]--
                if indegree[v] == 0 {
                    count++
                    q = append(q, v)
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