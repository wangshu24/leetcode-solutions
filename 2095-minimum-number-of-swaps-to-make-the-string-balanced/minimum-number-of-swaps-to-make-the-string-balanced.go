func minSwaps(s string) int {
    swap :=0
    //half := len(s)/2
    opening := 0
    for i:=0; i < len(s);i++{
        if string(s[i]) == "["  {opening++} else {
            if opening > 0 {opening--} else {
                fmt.Println(i)
                swap++
                opening++
            }
        }

    }

    return swap
}

// func mirror(s1, s2 string) bool {
//     if s1 == "[" && s2 == "]" {
//         return true
//     } else if s1 == "]" && s2 == "[" {
//         return true
//     } else {
//         return false
//         }
// }