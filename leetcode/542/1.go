func min(a,b int) int {
	if a <= b {
		return a
	}
	return b
}

func updateMatrix(martrix [][]int) [][]int {
	m,n := len(martrix), len(martrix[0])
	out := make([][]int, m)

	for i := 0; i < m; i++ {
		out[i] = make([]int,n)
		for  j := 0; j < n; j++ {
			if martrix[i][j] == 0 {
				out[i][j] = 0
			} else {
				out[i][j] = 10000
				if i > 0 {
					out[i][j] = min(out[i][j],out[i-1][j]+1)
				}
				if j > 0 {
					out[i][j] = min(out[i][j],out[i][j-1]+1)
				}
			}
		}
	}

	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if out[i][j] > 1 {
				if i < m-1 {
					out[i][j] = min(out[i][j], out[i+1][j]+1)
				}
				if j < n-1 {
					out[i][j] = min(out[i][j], out[i][j+1]+1)
				}
			}
		}
	}

}