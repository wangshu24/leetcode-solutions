func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1 = append(nums1[:m], nums2...)
    slices.Sort(nums1)

    //why does this not work???
    // for i:=0; i < n; i++ {
    //     nums1 = append(nums1, nums2[i])
    // }
    // nums1 = slices.Replace(nums1, m, m+n)
    //  fmt.Println(nums1)
    // slices.Sort(nums1)
    // fmt.Println(nums1)
}