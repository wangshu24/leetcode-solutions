func reverse(x int) int {
    num := strconv.Itoa(x)
    var b strings.Builder
    sign := 1
    for i:=len(num)-1; i >= 0; i-- {
        if num[i] == '-' {
            sign = -1
        } else {
            b.WriteByte(num[i])
        }
    }
    res, _ := strconv.Atoi(b.String())
    if res >= math.MaxInt32 {return 0}
    return res * sign
}