package caps

func genMatrix(i int) [][]int {
	m := make([][]int, i)

	for x := 0; x < i; x++ {
		m[x] = make([]int, i)
		for y := 0; y < i; y++ {
			m[x][y] = x*i + y
		}
	}

	return m
}
