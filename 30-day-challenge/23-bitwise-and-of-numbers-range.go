func rangeBitwiseAnd(m int, n int) int {
	res := m
	for i := m+1; i <= n; i++ {
		res = res & i
	}
	return res
}

// https://youtu.be/tM-FPbMh-SQ
func rangeBitwiseAnd2(m int, n int) int {
	var i int
	for m != n {
        // fmt.Printf("%b\n", m)
        // fmt.Printf("%b\n", n)
		m >>= 1
		n >>= 1
		i++
	}
	return n << i
}
