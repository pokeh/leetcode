func intervalIntersection(A [][]int, B [][]int) [][]int {
	ai, bi := 0, 0
	res := make([][]int, 0)

	for ai < len(A) && bi < len(B) {
		start := max(A[ai][0], B[bi][0])
		var end int

		if A[ai][1] < B[bi][1] {
			end = A[ai][1]
			ai++
		} else {
			end = B[bi][1]
			bi++
		}

		if start > end {
			continue
		}
		res = append(res, []int{start, end})
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
