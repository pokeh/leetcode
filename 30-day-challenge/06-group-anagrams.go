func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	tmp := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(a, b int) bool { return runes[a] < runes[b] })

		sorted := string(runes)
		tmp[sorted] = append(tmp[sorted], s)
	}

	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}

func groupAnagrams2(strs []string) [][]string {
	const first = rune('a')

	tmp := make(map[[26]int][]string)
	for _, s := range strs {
		key := [26]int{}
		for _, r := range s {
			key[r-first]++
		}
		tmp[key] = append(tmp[key], s)
	}

	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}
