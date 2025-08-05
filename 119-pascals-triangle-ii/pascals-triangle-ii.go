func getRow(row int) []int {
    res := [][]int{{1}}

    for i:=1; i <= row; i++ {
        n := make([]int, i+1)
        n[0], n[i] = 1, 1
        for j:=1; j < i; j++ {
            n[j] = res[i-1][j-1] + res[i-1][j]
        }
        res = append(res, n)
    }

    return res[row]
}