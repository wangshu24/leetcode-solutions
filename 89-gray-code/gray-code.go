func grayCode(n int) []int {
    var v []int
    for i:=0; i<1<<n; i++ {
        fmt.Printf("%0*b %0*b \n",n,  i^(i>>1), n,  i>>1)
        v =append(v, i^(i>>1));
    }
    return v;
}