func removeOccurrences(s string, part string) string {
    var builder strings.Builder
    var bpart =  bytes.NewBufferString(part).Bytes()
    tmp := make([]byte,0)
    
    for i:=0; i < len(s); i ++ {
        tmp = append(tmp, s[i])
        if len(tmp) >= len(part) {
            cur := tmp[len(tmp)-len(part):len(tmp)] 
            if  slices.Equal(cur, bpart) {
                tmp = tmp[:len(tmp)-len(part)]
            }
        }
    }
    
    builder.Write(tmp)
    return builder.String()
}