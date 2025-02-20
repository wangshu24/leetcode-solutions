func findDifferentBinaryString(nums []string) string {
    set := map[string]int{}
    res := make([]byte, len(nums))
    for i:=0; i < len(nums); i++ {
       set[nums[i]]++
    }
    var backtrack func(i int, cur []byte) string
    backtrack = func(i int, cur []byte) string {
        if i == len(nums) {
            if set[string(cur)] > 0 {return ""} else {
                return string(cur)
            }
        }
        for _,bit := range []byte{'0','1'} {
            cur[i] = bit
            newString := backtrack(i+1, cur)
            if string(newString) != "" {
                return string(newString)
            }
        }
        return ""
    }
    fmt.Println(backtrack(0, res))
    fmt.Println(res)
     
    return string(res)
}