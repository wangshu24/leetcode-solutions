func maxMatrixSum(mat [][]int) int64 {
    res:=0
    minVal := int(math.Pow(10,5))
    countNeg:= 0 
    for _,row := range mat {
        for ind,val := range row {
            if val <= 0 {
                countNeg++
                row[ind] = val * -1
                minVal = min(minVal, row[ind])
                res = res + row[ind]
            } else {
                minVal = min(minVal, val)
                res = res + val
            }
        }      
    }

    if countNeg%2 != 0 {
        res = res - (minVal*2)
    }

    return int64(res)
}

func min (a,b int) int {
    if a < b {
        return a
    }
    return b
}