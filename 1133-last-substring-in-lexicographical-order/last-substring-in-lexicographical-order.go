func lastSubstring(s string) string {
    max := len(s)-1
    for cur:= len(s)-1; cur >= 0;cur-- {
        if s[cur] > s[max]{
            max = cur
        } else if s[cur] == s[max] {
            i,j := cur+1, max+1

            for i < max && j < len(s) && s[i] == s[j] {
                i++
                j++
            }

            if  i == max || j == len(s) || s[i] > s[j] {
                max = cur
            }
        }
    }
    fmt.Println(max)
    return s[max:]
}   