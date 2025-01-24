func countServers(grid [][]int) int {
    row := make([]int, len(grid))
    col := make([]int, len(grid[0]))

    for i:= 0; i < len(grid); i++ {
        for j:= 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                row[i]++
                col[j]++
            }
        }
    }

    result:=0
    for i:=0; i < len(grid); i++ {
        for j:=0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 && ( row[i] > 1 || col[j] > 1 ) {
                result++
            }        
        }
    }


    return result 
}