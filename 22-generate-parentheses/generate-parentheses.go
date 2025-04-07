func generateParenthesis(n int) []string {
    var res []string

    var backtrack func(i, n int, subset []byte)

    backtrack = func(i, n int, subset []byte) {
        if i == n {
            if wellform(string(subset)) {
                res = append(res, string(subset))
            }
            return
        } else if i > n {return}

        subset = append(subset, '(')
        backtrack(i+1, n, subset)
        subset = subset[:len(subset)-1]
        subset = append(subset, ')')
        backtrack(i+1, n, subset)
    }

    backtrack(0, n+n, make([]byte,0))
    return res
}

func wellform(s string) bool {
    open := []byte{}
    for _,char := range s {
        if char == '(' {
            open = append(open, '(')
        } else {
            if len(open) == 0 {
                return false
            } else if open[len(open)-1] == '(' {
                open = open[:len(open)-1]
            }
        }
    }
    return len(open) == 0
}