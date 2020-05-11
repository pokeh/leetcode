// Runtime: 104 ms
// Memory Usage: 7.1 MB
func findJudge(N int, trust [][]int) int {
	trustIndex := make([]int, N)

	for _, t := range trust {
		trustIndex[t[0]-1]--
		trustIndex[t[1]-1]++
	}

	for i, val := range trustIndex {
		// trusted by everyone else && does not trust anyone
		if val == N-1 {
			return i + 1
		}
	}

	return -1
}

/*
// Runtime: 204 ms
// Memory Usage: 7.4 MB
func findJudge(N int, trust [][]int) int {
	trusteeCnt := make(map[int]int)
	trustsNoOne := make(map[int]struct{})

	for i := 1; i <= N; i++ {
		trustsNoOne[i] = struct{}{}
	}

	for _, t := range trust {
		trusteeCnt[t[1]]++

		if _, ok := trustsNoOne[t[0]]; ok {
			delete(trustsNoOne, t[0])
		}
	}

	if len(trustsNoOne) != 1 {
		return -1
	}
	for candidate, _ := range trustsNoOne {
		if trustCnt[candidate] == N-1 {
			return candidate
		}
	}

	return -1
}
*/
