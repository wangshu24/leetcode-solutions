func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i:=0; i < n; i++ {
        fmt.Println(len(nums1))
        nums1 = slices.Replace(nums1, m+i, m+i+1, nums2[i])
    }
    slices.Sort(nums1)
    // for i:=0; i < n; i++ {
    //     nums1 = append(nums1, nums2[i])
    // }
    // nums1 = slices.Replace(nums1, m, m+n)
    //  fmt.Println(nums1)
    // slices.Sort(nums1)
    // fmt.Println(nums1)
}