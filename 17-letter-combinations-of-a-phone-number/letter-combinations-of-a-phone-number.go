func letterCombinations(digits string) []string {
    var res []string
    if len(digits) == 0 {return res}
    
    l := map[byte]string{
        '2' : "abc", 
        '3' : "def",
        '4' : "ghi",
        '5' : "jkl",
        '6' : "mno",
        '7' : "pqrs",
        '8' : "tuv",
        '9' : "wxyz",
        }
    
    var backtrack func(i int,subset []byte)

    backtrack = func(i int, subset []byte) {
        if i == len(digits) {
            if len(subset) == len(digits) {
                res = append(res, string(subset))
            }
            return
        } else if i > len(digits) {return}

        chars := l[digits[i]]
        
        for _, char := range chars {
            subset = append(subset, byte(char))
            backtrack(i+1, subset)
            subset = subset[:len(subset)-1]
        } 
    }

    backtrack(0, make([]byte,0))
    return res
}