func canArrange(arr []int, k int) bool {
    remain := map[int]int{}
    for i:=0; i< len(arr); i++ {
        remain[((arr[i]%k)+k)%k]++
    }
    fmt.Println(remain)

    for key,val := range remain {
        if key == 0 { if val%2 != 0 {
            return false
            }
        } else {
            if remain[k-key] != val {return false}
        }
    }

    return true
}









// func canArrange(arr []int, k int) bool {
//     slices.Sort(arr)
//     largeArr := []int{}
//     smallArr  := []int{}
//     result := true
//     for i:=len(arr)-1;i >-1;i-- {
//         if arr[i]>k {largeArr = append(largeArr, arr[i])} else {smallArr = append(smallArr, arr[i])}
//     }
  
//     fmt.Println(smallArr)
//     fmt.Println(largeArr)
    

//     for i:=0; i<len(largeArr); i++ {
//         fmt.Println("got here")
//         if largeArr[i]%k == 0 {
//             if !slices.ContainsFunc(smallArr, func(n int) bool {return n%k == 0}) {return  false}
//         } else {
//             if !slices.ContainsFunc(smallArr, func(n int) bool {return n == largeArr[i]%k}) {return false}
//         }
//     } 
//     if len(largeArr) == 0 {
//         for i:=0;i < len(smallArr); i++ {
//             if !slices.ContainsFunc(smallArr, func(n int) bool {return n == (k - smallArr[i])}) {return false}
//         }
//     }

//     return result
// }

