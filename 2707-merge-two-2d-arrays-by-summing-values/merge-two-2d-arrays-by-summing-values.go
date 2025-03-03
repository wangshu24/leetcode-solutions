func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    // Initialize pointers for both arrays
    i, j := 0, 0
    result := [][]int{}
    
    // Traverse both arrays with two pointers
    for i < len(nums1) && j < len(nums2) {
        id1, val1 := nums1[i][0], nums1[i][1]
        id2, val2 := nums2[j][0], nums2[j][1]
        
        if id1 < id2 {
            // Id1 is smaller, add it to result
            result = append(result, []int{id1, val1})
            i++
        } else if id2 < id1 {
            // Id2 is smaller, add it to result
            result = append(result, []int{id2, val2})
            j++
        } else {
            // Ids are equal, sum the values
            result = append(result, []int{id1, val1 + val2})
            i++
            j++
        }
    }
    
    // Add remaining elements from nums1, if any
    for i < len(nums1) {
        result = append(result, nums1[i])
        i++
    }
    
    // Add remaining elements from nums2, if any
    for j < len(nums2) {
        result = append(result, nums2[j])
        j++
    }
    
    return result
}