func maxWidthRamp(nums []int) int {
    dai := len(nums)
    stacks := []int{}
    for i:=0; i< dai;i++{
        if len(stacks) ==0 || nums[stacks[len(stacks)-1]] > nums[i]{
            stacks = append(stacks, i)
        }
    }

    maxWid := 0
    for i:=dai-1; i >=0; i--{
        for len(stacks)>0 && nums[stacks[len(stacks)-1]] <= nums[i]{
            j := stacks[len(stacks)-1]
            stacks = stacks[:len(stacks)-1]
            if i - j > maxWid{
                maxWid = i - j
            }
        }
    }

    return maxWid
}