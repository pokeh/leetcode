func maxSubarraySumCircular(A []int) int {
	globalMax, localMax := A[0], A[0]
	globalMin, localMin := A[0], A[0]
	sum := A[0]

	for i := 1; i < len(A); i++ {
		localMax = max(A[i], localMax+A[i])
		globalMax = max(globalMax, localMax)

		localMin = min(A[i], localMin+A[i])
		globalMin = min(globalMin, localMin)

		sum += A[i]
	}

	max1 := globalMax
	max2 := sum - globalMin

	if max1 > max2 || max2 == 0 {
		return max1
	}
	return max2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
// O(N^2)
// Time limit exceeded
func maxSubarraySumCircular(A []int) int {
	globalMax := A[0]

	for i := 0; i < len(A); i++ {
		iLocalMax := A[i]
		iGlobalMax := A[i]

		for j := i + 1; j < len(A)+i; j++ {
			index := j % len(A)
			iLocalMax = max(iLocalMax+A[index], A[index])
			iGlobalMax = max(iLocalMax, iGlobalMax)
		}

		globalMax = max(iGlobalMax, globalMax)
	}

	return globalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/
