// func maxArea(height []int) int {
//     result :=1
//     highest, maxVolInd := 0, 0 
//     for ind,val := range height {
//         if val > highest {
//             highest = val
//         }
        
//         newVol := 1
//         if val > highest {
//             newVol = highest * (ind - maxVolInd) 
//         } else {
//             newVol = val * (ind - maxVolInd)
//         }
//         if newVol > result {
//             maxVolInd = ind
//             }
//         result = max(result, newVol)
//         fmt.Println(result)
//     }
//     return result
// }

func maxArea(height []int) int {
    res := 0
    left, right := 0, len(height)-1 

    for left < right {
        valLeft := height[left]
        valRight := height[right]
        tempMax := int(math.Min(float64(valLeft), float64(valRight))) * (right - left)
        res = max(res, tempMax)

        if  valLeft < valRight {
            left++
        } else {
            right --
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
        }
}