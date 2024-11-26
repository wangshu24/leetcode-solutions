func findChampion(n int, edges [][]int) int {
    if len(edges) == 0  {
        if n == 1 {
            return 0
        } else {
            return -1
        }
    }
    team := map[int]int{}
    
    for _,ed := range edges {
        if team[ed[0]] == 0 { team[ed[0]] = 0 }
        team[ed[1]]++
    }

    res:= []int{}
    teamFound := 0
    for key,val := range team{
        teamFound++
        if val == 0 {
            res = append(res, key)
        }
    }
    if teamFound < n {return -1}
    if len(res) > 1 {return -1}

    return res[0]
}   