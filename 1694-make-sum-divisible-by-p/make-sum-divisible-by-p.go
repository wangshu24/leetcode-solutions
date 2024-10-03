/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
func minSubarray(nums []int ,p int ) int {
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

    for i:=0; i< len(nums);i++{
        prefixSum+= nums[i]
        currentRem := prefixSum%p
        targetRem := (currentRem - rem + p)%p //we add p then divide remainder against p to ensure the gap between current remain and final remain is positive

        if id, ok := remainMap[targetRem]; ok {
            if i - id < minLength {
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


