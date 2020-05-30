func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	// directed graph of edges prereq -> course
	graph := make(map[int][]int, numCourses)
	for _, p := range prerequisites {
		course, prereq := p[0], p[1]
		graph[prereq] = append(graph[prereq], course)
	}

	// canFinish <=> cycle doesn't exist in the directed graph
	explored := make([]int, numCourses)

	var hasCycle func(graph map[int][]int, vertex int) bool
	hasCycle = func(graph map[int][]int, vertex int) bool {
		if explored[vertex] == 1 {
			return true
		}
		if explored[vertex] == 2 {
			return false
		}

		explored[vertex] = 1 // start exploring
		if outNeighbours, ok := graph[vertex]; ok {
			for _, n := range outNeighbours {
				if hasCycle(graph, n) {
					return true
				}
			}
		}
		explored[vertex] = 2 // finish exploring
		return false
	}

	for i := 0; i < numCourses; i++ {
		if explored[i] == 0 {
			if hasCycle(graph, i) {
				return false
			}
		}
	}

	return true
}
