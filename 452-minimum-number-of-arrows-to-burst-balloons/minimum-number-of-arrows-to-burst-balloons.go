import (
    "slices" 
    "cmp"
)

func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(a, b []int) int {
            return cmp.Compare(a[1], b[1])
    })
    fmt.Println(points)

    res, minx := 1, points[0][1]
    for _, point := range points {
        if minx >= point[0] {
            continue
        } else {
            minx = point[1]
            res++
        }
    }

    return res
}