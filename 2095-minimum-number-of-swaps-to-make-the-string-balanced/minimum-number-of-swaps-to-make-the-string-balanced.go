func minSwaps(s string) int {
    swap :=0
    opening := 0
    for i:=0; i < len(s);i++{
        if string(s[i]) == "["  {opening++} else {
            if opening > 0 {opening--} else {
                swap++
                opening++
            }
        }

    }

    return swap
}