func majorityElement(nums []int) int {
    hashmap := map[int]int{}
    for i:=0; i < len(nums); i++ {
        hashmap[nums[i]]++
    }
    count := 0
    output :=0
    for i, j := range hashmap {
        fmt.Println("key: ", i, " value is: ", j)
        if(j > count) {count = j; output = i}
    }
    return output
}