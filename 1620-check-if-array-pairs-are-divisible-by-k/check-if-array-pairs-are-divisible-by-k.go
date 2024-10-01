func canArrange(arr []int, k int) bool {
    remain := map[int]int{}
    for i:=0; i< len(arr); i++ {
        remain[((arr[i]%k)+k)%k]++
    }

    for key,val := range remain {
        if key == 0 { if val%2 != 0 {
            return false
            }
        } else {
            if remain[k-key] != val {return false}
        }
    }
    return true
}

