package _022_07_30

func convert(s string, numRows int) string {
	r := numRows
	if r < 2 {
		return s
	}
	n := len(s)
	t := 2*r - 2
	tn := (n + t - 1) / t * (r - 1)
	table := make([][]byte, r)
	for i := 0; i < r; i++ {
		table[i] = make([]byte, tn)
	}
	x, y := 0, 0
	for i, ch := range s {
		table[x][y] = byte(ch)
		if i%t < r-1 {
			x++
		} else {
			x--
			y++
		}
	}
	res := make([]byte, 0)
	for i := 0; i < r; i++ {
		for j := 0; j < tn; j++ {
			if table[i][j] > 0 {
				res = append(res, table[i][j])
			}
		}
	}
	return string(res)
}
