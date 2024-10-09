func minAddToMakeValid(s string) int {
    if len(s) == 0 {return 0}
    opening :=0
    add:=0
    for i:=0; i < len(s);i++{
        if string(s[i]) == "(" {
            opening++
            add++
        } else{
            if opening > 0 {
                opening--
                add--
                } else {
                    // opening++
                    add++
                } 
        }
    }
    return add
}