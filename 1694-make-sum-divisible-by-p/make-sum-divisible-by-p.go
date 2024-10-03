/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
func minSubarray(nums []int ,p int ) int {
    //sub := []int{}
    numsSum := 0
    remainMap := map[int]int{}
    for i:=0; i <len(nums);i++ {
        numsSum+=nums[i]    
    }
    
    if numsSum < p {return -1}
    rem:= numsSum%p
    remainMap[0] = -1
    prefixSum, minLength := 0, len(nums)

     if numsSum%p == 0 {
        return 0
    } 

    // if remain%p == 0 {
    //     return 0
    // }  else if slices.Contains(remainMap, remain) { return 1} else {
    //     for i:=0; i<len(remainMap);i++{
    //         remain-=remainMap[i]
    //         fmt.Println("remainder to be subtracted ", remainMap[i], " remain = ", remain)
    //         if remain == 0 {
    //             sub = append(sub,remainMap[i] )
    //             return len(sub)
    //         } else if remain > 0 {
    //             sub = append(sub,remainMap[i] )
    //         } else {
    //             remain = numsSum%p
    //             sub = nil
    //         }
    //     }
    //     fmt.Println("final remain is: ", remain)
    //     if remain != 0 {return -1}
    // }
    
    for i:=0; i< len(nums);i++{
        prefixSum+= nums[i]
        currentRem := prefixSum%p
        targetRem := (currentRem - rem + p)%p //we add p then divide remainder against p to ensure the gap between current remain and final remain is positive

        if id, ok := remainMap[targetRem]; ok {
            if i - id < minLength {
                fmt.Println("got here")
                minLength = i - id
            }
        }
        remainMap[currentRem] = i
    }
    if minLength == len(nums){
        return -1
    }

    return minLength
};


