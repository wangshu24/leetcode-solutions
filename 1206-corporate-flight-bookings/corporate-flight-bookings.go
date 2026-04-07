func corpFlightBookings(bookings [][]int, n int) []int {
    var last int
    for _,val := range bookings {
        last = max(last, val[1])
    } 
    resLen := max(last, len(bookings))
    res := make([]int, resLen)
    for _, val := range bookings {
        s, e := val[0], val[1]
        res[s-1]+=val[2]
        if e < len(res) {res[e]-=val[2]}
    }

    for i:=1; i < len(res); i++ {
        res[i] = res[i-1] + res[i]
    }
    fmt.Println(res)
    return res
}