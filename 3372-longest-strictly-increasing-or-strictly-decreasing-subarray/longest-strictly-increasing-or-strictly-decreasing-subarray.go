func max(a, b int) int {
    if a > b {return a}
    return b
}

// func longestMonotonicSubarray(nums []int) int {
//     temp, res, root := 1, 1, nums[0]

//     for i:= 1; i < len(nums); i++ {
//         if nums[i] > root  && nums[i] > nums[i-1] {
//             temp++
//         } else if nums[i] < root && nums[i] < nums[i-1] {
//             temp++
//         } else {
//             res  = max(res, temp)
//             temp = 1
//             root = nums[i]
//         }
//         fmt.Println(temp, res)

//     } 
    
//        fmt.Println(temp, res)
//     res = max(res, temp)
//     return res
// }

func longestMonotonicSubarray(nums []int) int {
    temp := []int{}
    res, tren:=1, 0
    // 1 ascend, 0 descent

    for i:=0; i < len(nums); i++ {
        if len(temp) == 0 {
            temp = append(temp, nums[i])
        } else {
            if len(temp) == 1 {
                if nums[i] > temp[0] {
                    tren = 1
                    temp = append(temp, nums[i])
                } else if nums[i] < temp[0] {
                    tren = 0
                    temp = append(temp, nums[i])
                } else {
                    temp[0] = nums[i]
                }
            } else {
                if tren == 0 {
                    if nums[i] < temp[len(temp)-1] {
                        temp = append(temp, nums[i])
                    } else {
                        res = max(res, len(temp))
                        if nums[i] == temp[len(temp)-1] {
                            temp = []int{ nums[i]}
                           
                        } else {
                            temp = []int{nums[i-1], nums[i]}
                            tren = 1
                        }
                        
                    }
                } else {
                    if nums[i] > temp[len(temp)-1] {
                        temp = append(temp, nums[i])
                    } else {
                        res = max(res, len(temp))
                        if nums[i] == temp[len(temp)-1] {
                            temp = []int{ nums[i]}
                        } else {
                            temp = []int{nums[i-1], nums[i]}
                            tren = 0
                        }
                    }
                }
            }
        }
        fmt.Println(temp, tren)
    }
    res = max(res, len(temp))
    return res
}
