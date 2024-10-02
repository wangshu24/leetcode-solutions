func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    for i:=0; i< len(nums2); i++ {
        nums1 = append(nums1, nums2[i])
    }
    slices.Sort(nums1)
    var result float64
    if len(nums1)%2 == 1 {
        result = float64(nums1[(len(nums1)-1)/2])
    } else {
        medianS := float64(nums1[len(nums1)/2 -1])
        medianL := float64(nums1[len(nums1)/2 ])
        result = (medianS + medianL) /2
    }
    return result
}