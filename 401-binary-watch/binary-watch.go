import "strconv"

func readBinaryWatch(turnedOn int) []string {
    d := []int{8,4,2,1,32,16,8,4,2,1}
    res := []string{}

    var dfs func(turnedOn, H, M, k int)
    dfs = func(turnedOn, H, M, k int) {
        if turnedOn == 0 {
            h, m := strconv.Itoa(H), strconv.Itoa(M)
            if M < 10 {
                m = "0" + m
            }
            res = append(res, h + ":" + m)
            return
        }

        for i:=k; i < 10;i++{
            h, m := H, M

            if i < 4 {
                h += d[i]
            } else {
                m += d[i]
            }

            if h < 12 && m < 60 {
                dfs(turnedOn-1, h, m, i+1)
            }
        }
    }

    dfs(turnedOn, 0,0,0)

    return res
}

