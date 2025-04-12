func exist(board [][]byte, word string) bool {
    var (
        ROW = len(board)
        COL = len(board[0])
        visit = make(map[string]bool)
    )
    var backtrack func(i, x, y int) bool

    backtrack = func(i, x, y int) bool {
        if i == len(word) {
            return true
        } else if x >= ROW || y >= COL || x < 0 || y < 0 || 
        board[x][y] != word[i] || visit[strconv.Itoa(x)+","+strconv.Itoa(y)] {
            return false
        }

        visit[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
        res := backtrack(i+1, x+1,y) ||
        backtrack(i+1, x-1,y) ||
        backtrack(i+1, x,y+1) ||
        backtrack(i+1, x,y-1)
        visit[strconv.Itoa(x)+","+strconv.Itoa(y)] = false
        return res 
    }

    for i:=0; i < ROW; i++ {
        for j:=0; j < COL; j++ {
            if backtrack(0,i,j) {return true}
        }
    }

    return false
}