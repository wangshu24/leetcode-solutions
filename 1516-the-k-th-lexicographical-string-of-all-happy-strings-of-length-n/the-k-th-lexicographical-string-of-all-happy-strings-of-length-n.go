func getHappyString(n int, k int) string {
    var res string
    var strs []string
    chars := []byte{ 'a','b','c'}
    
    var createString func(int, []byte, byte)
    createString = func (strlen int, str []byte, prev byte) {
        if strlen == n {
            cp := string(str)
            strs = append(strs, cp)
            return
        }
        for i:=0; i < 3; i++ {
            if  prev != chars[i] {
                str = append(str, chars[i])
                createString(strlen+1, str, chars[i])
                str = str[:len(str)-1]
            }
        } 
        return 
    }

    createString(0, []byte{}, '0')
    
    for len(strs) < k {
        return ""
    }

    res = strs[k-1]
    return res 
}