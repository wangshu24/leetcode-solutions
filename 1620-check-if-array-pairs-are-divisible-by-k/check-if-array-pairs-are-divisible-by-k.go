func canArrange(arr []int, k int) bool {
    // Frequency array to store the count of remainders
    remainderFreq := make([]int, k)
    
    // Step 1: Calculate the remainder for each element and store the frequency
    for _, num := range arr {
        remainder := ((num % k) + k) % k // Ensure non-negative remainder
        remainderFreq[remainder]++
    }
    
    // Step 2: Check if the pairing condition holds
    for i := 0; i <= k/2; i++ {
        if i == 0 {
            // Elements with remainder 0 must pair among themselves
            if remainderFreq[i] % 2 != 0 {
                return false
            }
        } else {
            // Remainder i must pair with remainder k-i
            if remainderFreq[i] != remainderFreq[k-i] {
                return false
            }
        }
    }
    
    return true
}

// func canArrange(arr []int, k int) bool {
//     remain := map[int]int{}
//     for i:=0; i< len(arr); i++ {
//         remain[arr[i]] = arr[i]%k
//     }
//     fmt.Println(remain)

//     for key,val  := range remain {
//         if key < k {
//             if !slices.ContainsFunc(arr, func(n int) bool {return (n+key)%k == 0}) { 
//                 fmt.Println("got ", key, " lower")
//                 return false}
//         } else {
//             if !slices.ContainsFunc(arr, func(n int) bool {return n == val || n%k==val}) { 
//                 fmt.Println("got ", key, " higher")
//                 return false}
//         }
//     }

//     return true
// }









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

