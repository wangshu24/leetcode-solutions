func shortestCommonSupersequence(str1 string, str2 string) string {
    len1, len2 := len(str1), len(str2)
    prev := make([]string, len1+1)
    for i:=0 ; i <= len1; i++ {
        prev[i] = str1[:i]
    }
    
    for row:=1; row <= len2; row++ {
        cur := make([]string, len1+1)
        cur[0] = str2[:row]
        for col:=1; col <= len1; col++ {
            if str1[col-1] == str2[row-1] {
                cur[col] = prev[col-1] + string(str2[row-1])
            } else {
                res1 := prev[col]
                res2 := cur[col-1]
                if len(res1) < len(res2) {
                    cur[col] = res1 + string(str2[row-1])
                } else {
                    cur[col] = res2 + string(str1[col-1])
                }
            }
        }
        prev = cur
    }

    return prev[len1]
}