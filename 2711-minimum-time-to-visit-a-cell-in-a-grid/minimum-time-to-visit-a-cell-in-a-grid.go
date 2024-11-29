type Cell struct {
  tyme, row, col int
}

type PriorQ []*Cell

func (pq PriorQ) Len() int { return len(pq) }
func (pq PriorQ) Less(i, j int) bool { return pq[i].tyme < pq[j].tyme }
func (pq PriorQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorQ) Push(x any) { *pq = append(*pq, x.(*Cell)) }
func (pq *PriorQ) Pop() any {
	old := *pq
	n := len(old)
	cell := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return cell
}

func minimumTime(G [][]int) int {
  if G[0][1] > 1 && G[1][0] > 1 { return -1 }
  row, col := len(G), len(G[0])
  dirs := [4][2]int{{1,0}, {0,1}, {-1,0}, {0,-1}}
  V := make([][]bool, row)
  for i := range V {
    V[i] = make([]bool, col)
  }
  V[0][0] = true

  Q := &PriorQ{}
  heap.Push(Q, &Cell{0, 0, 0})

  for Q.Len() > 0 {
    cell := heap.Pop(Q).(*Cell)

    for _, d := range dirs {
      r, c := cell.row+d[0], cell.col+d[1]
      if 0 <= r && r < row && 0 <= c && c < col && !V[r][c] {
        t := cell.tyme + 1
        if t < G[r][c] {
          t = G[r][c] + (G[r][c] - t) % 2
        }

        if r == row-1 && c == col-1 { return t }
        heap.Push(Q, &Cell{t, r, c})
        V[r][c] = true
      }
    }
  }

  return -1
}