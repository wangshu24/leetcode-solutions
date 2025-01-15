func canBeValid(s string, locked string) bool {
    if len(s) <2 || len(s)%2 != 0 {return false}
    if !strings.Contains(locked, "1") {return true}
   
    lb, rb := 0,0

    for l:=0; l < len(s); l++ {
        if locked[l] == 48 || s[l] == 40 {
            lb++
        } else if s[l] == 41 {
            lb--
        }
        if lb < 0 {return false}
    }

    for r:=len(s)-1; r > 0; r-- {
        if locked[r] == 48 || s[r] == 41 {
            rb++
        } else if s[r] == 40 {
            rb--
        }
        if rb < 0 {return false}
    }

    return true
} 