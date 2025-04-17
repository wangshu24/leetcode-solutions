func restoreIpAddresses(s string) []string {
    var res []string

    var backtrack func(i, dot int, set []byte)
    backtrack = func(i, dot int, set []byte) {
        if i == len(s) {
            tmp := string(set)
            if valid(tmp) {
                res = append(res, tmp)
            }
            return
        }

        set = append(set, s[i])
        backtrack(i+1, dot,set)
        if dot > 0 {
            set = append(set, '.')
            backtrack(i+1, dot-1, set)
        }
    }
    backtrack(0,3, make([]byte, 0))

    return res
}

func valid(s string) bool {
    set := strings.Split(s, ".")
    if len(set) < 4 {return false}
    for _,str := range set {
        // if str == "" {return false}
        if len(str) == 0 || len(str) > 3 {return false}
        if len(str) > 1 && str[0] == '0' {return false}
        num,_ := strconv.Atoi(str)
        if num > 255 ||  num < 0 {return false}
    }
    return true
}