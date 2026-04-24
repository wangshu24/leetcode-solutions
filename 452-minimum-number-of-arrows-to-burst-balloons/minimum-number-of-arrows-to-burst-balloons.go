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
    for i:=0; i < len(points); i++  {
        if minx >= points[i][0] {
            continue
        } else {
            minx = points[i][1]
            res++
        }
    }

    return res
}