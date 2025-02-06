func tupleSameProduct(nums []int) int {
    x := make(map[int]int)
    
    for j:=0; j < len(nums); j++ {
        for i:=j+1; i < len(nums); i++ {
                x[nums[j]*nums[i]]++
        }
    }
    count :=0
    for _,val := range x {
        pair := ((val*(val-1))/2)
        count += 8*pair
    }

    return count
}