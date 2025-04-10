func grayCode(n int) []int {
    v := make([]int, 1 << n)
    for i:=0; i<1<<n; i++ {
        v[i] =  i^(i>>1)
    }
    return v;
}