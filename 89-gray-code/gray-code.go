func grayCode(n int) []int {
    if n == 0 {return []int{0}}

    if n == 1 {return []int{0,1}}
    
    s1 := grayCode(n-1)

    s2 := make([]int, len(s1) * 2)

    copy(s2, s1)
    for i := len(s1); i < len(s1) * 2; i++ {
        s2[i] = i ^ (i >> 1)
    }
    
    return s2
}

// func grayCode(n int) []int {
//     res := make([]int, 1 << n)

//     for i:=0; i < 1 << n; i++ {
//         res[i] = i ^ (i >> 1)
//     }

//     return res
// }