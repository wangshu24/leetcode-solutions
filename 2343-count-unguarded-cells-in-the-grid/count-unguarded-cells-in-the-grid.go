func countUnguarded(m int, n int, G [][]int, W [][]int) int {
  result := 0
  D := make([]byte, m*n)
  for _, w := range W {
    D[w[0]*n + w[1]] = 1
  }
  for _, g := range G {
    D[g[0]*n + g[1]] = 1
  }

  for _, g := range G {
    r, c := g[0], g[1]
    for r := r-1; r >= 0 && D[r*n+c] != 1; r-- { D[r*n+c] = 2 }
    for r := r+1; r < m && D[r*n+c] != 1; r++ { D[r*n+c] = 2 }
    for c := c-1; c >= 0 && D[r*n+c] != 1; c-- { D[r*n+c] = 2 }
    for c := c+1; c < n && D[r*n+c] != 1; c++ { D[r*n+c] = 2 }
  }

  for i := 0; i < m*n; i++ {
    if D[i] == 0 {
      result++
    }
  }
    
  return result
}