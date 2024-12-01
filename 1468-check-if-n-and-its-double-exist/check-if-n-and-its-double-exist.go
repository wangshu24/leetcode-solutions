func checkIfExist(arr []int) bool {
    sort.Ints(arr)
    fmt.Println(arr)
    for i:=0; i < len(arr);i++{
        l,r := 0, len(arr)-1
        x := 2 * arr[i]
        for l<=r {
            mid := (l+r)/2
            if arr[mid] == x && mid != i {
                return true
            } else if arr[mid] <x {
                l = mid +1
            } else {r = mid -1}
        }
    } 

    return false
}