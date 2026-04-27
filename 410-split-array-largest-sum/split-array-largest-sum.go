func splitArray(nums []int, k int) int {
    minsub := slices.Max(nums)
    var maxsub int 
    for _, val := range nums {
        maxsub+=val
    }

    for minsub < maxsub {
        mid := minsub + (maxsub - minsub)/2
        split, sum := 1, 0
        for i:=0; i < len(nums);i++ {
            if sum+nums[i] > mid {
                split++
                sum = nums[i]
            } else {
                sum+= nums[i]
            }
        }
        if split > k {
            minsub = mid+1
        } else {
            maxsub = mid
        }
    }


    return minsub
}

