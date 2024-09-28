func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {return ""}

    prefix := ""

    slices.Sort(strs)
    
    short := strs[0]
    long := strs[len(strs)-1]
    for i:=0; i < len(short); i++ {
        if string(short[i]) != string(long[i]){
            return prefix 
        } else {
            prefix += string(short[i])
        }
    }

    return prefix
}