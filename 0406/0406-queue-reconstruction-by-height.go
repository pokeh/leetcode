func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	res := make([][]int, 0, len(people))

	for _, p := range people {
		// insert into k-th position
		k := p[1]
		tmp := append([][]int{}, res[k:]...)
		res = append(append(res[:k], p), tmp...)
	}

	return res
}
