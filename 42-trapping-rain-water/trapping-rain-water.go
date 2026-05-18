func trap(height []int) int {
    q := make([]int,0)
    total := 0

    for i:=0; i < len(height); i ++ {
        for len(q) > 0 && height[i] > height[q[len(q)-1]] {
            bId := q[len(q)-1]
            q = q[:len(q)-1]
            if len(q) == 0 {
                break
            }
            lW := q[len(q)-1]
            minH := min(height[i], height[lW]) - height[bId]
            w := i - lW -1
            total += minH * w
        }
        q = append(q, i)
    }

    return total
}