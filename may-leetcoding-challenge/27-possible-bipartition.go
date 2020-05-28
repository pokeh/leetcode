// tactic: create a graph with edges between the dislikes
// then try to "color the graph by 2 colors", i.e. start off with one vertex as 1,
// then assign -1 to its adjacent edges, then 1 to the next adjacent... etc
// if it contradicts at some point, return false
func possibleBipartition(N int, dislikes [][]int) bool {
	if len(dislikes) == 0 {
		return true
	}

	// adjacent edges
	adjLists := make(map[int][]int, N)

	for _, d := range dislikes {
		a, b := d[0], d[1]
		adjLists[a] = append(adjLists[a], b)
		adjLists[b] = append(adjLists[b], a)
	}

	// fmt.Println(adjLists)

	// group is either -1 or 1
	// length is N+1 so indices can match the edge numbers (we don't use the 0th element)
	groups := make([]int, N+1)

	for i := 1; i <= N; i++ {
		if groups[i] == 0 && !dfs(i, 1, adjLists, groups) {
			return false
		}
	}

	return true
}

func dfs(edge int, group int, adjLists map[int][]int, groups []int) bool {
	if groups[edge] != 0 {
		return groups[edge] == group
	}
	groups[edge] = group

	for _, adj := range adjLists[edge] {
		if !dfs(adj, -1*group, adjLists, groups) {
			return false
		}
	}

	return true
}

/*
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
*/
