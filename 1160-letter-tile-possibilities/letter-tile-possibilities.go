func numTilePossibilities(tiles string) int {
    dist := make(map[byte]int)
    for i:=0; i < len(tiles); i++ {
        dist[tiles[i]]++
    }

    var calc func() int
    calc = func() int {
        var res int
        for key,_ := range dist {
            if dist[key] > 0 {
                dist[key]--
                res++
                res += calc()
                dist[key]++
            }
        }
        return res
    }

    return calc()
}

