// func areAlmostEqual(s1 string, s2 string) bool {
//     if s1 == s2 {return true}
//     if len(s1) != len(s2) {return false}
    
//     mismatch := 0
//     for i:=0; i < len(s1); i++ {
//         if s1[i] != s2[i] {mismatch++}
//     }
//     if mismatch == len(s1) {return false}
//     if mismatch == 2 {
//         return true
//     }

//     return false
// }

func areAlmostEqual(s1 string, s2 string) bool {
    if s1 == s2 {return true}
    if len(s1) != len(s2) {return false}
    l:=0
    chars1 := map[byte]int{}
    chars2 := map[byte]int{}
    for i:=0; i < len(s1) ; i++{
        if s1[i] != s2[i] {
            chars1[s1[i]]++
            chars2[s2[i]]++
            l++
        }
    }
    if l > 2 {return false}
    if maps.Equal(chars1, chars2) {
        return true
    }
    return false
}