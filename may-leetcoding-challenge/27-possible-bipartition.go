// thanks to https://leetcode.com/problems/possible-bipartition/discuss/543270/golang
func possibleBipartition(N int, dislikes [][]int) bool {
	if len(dislikes) == 0 {
		return true
	}

	// graph[a][b] == true means that there is an edge between a and b
	// create N+1 so indices match with edge numbers (we don't use the 0th element)
	graph := make([][]bool, N+1)
	for i := 1; i <= N; i++ {
		graph[i] = make([]bool, N+1)
	}
	for _, dislike := range dislikes {
		graph[dislike[0]][dislike[1]] = true
		graph[dislike[1]][dislike[0]] = true
	}

	// colors are either -1 or 1
	colors := make([]int, N+1)

	// try coloring the first edge as 1,
	// then its adjecent edges as -1, then the adjecent edges of those as 1, etc...
	// if the color(-1/1) of an edge doesn't match what's expected, return false
	for i := 1; i <= N; i++ {
		if colors[i] == 0 {
			if !dfs(i, N, 1, colors, graph) {
				return false
			}
		}
	}

	return true
}

func dfs(i, N, expect int, colors []int, graph [][]bool) bool {
	if colors[i] != 0 {
		return colors[i] == expect
	}
	colors[i] = expect

	for j := 1; j <= N; j++ {
		if graph[i][j] && !dfs(j, N, -1*expect, colors, graph) {
			return false
		}
	}

	return true
}
