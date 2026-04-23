func corpFlightBookings(bookings [][]int, n int) []int {
    res := make([]int, n+1)

    for _, book := range bookings {
        res[book[0]-1] += book[2]
        res[book[1]] -= book[2]
    }
    fmt.Println(res)
    for i:=1; i < len(res); i++ {
        res[i] += res[i-1]
    }
    fmt.Println(res)

    return res[:n]
}