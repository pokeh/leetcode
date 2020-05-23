type charCount struct {
	char  rune
	count int
}

func frequencySort(s string) string {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}

	var charCounts []charCount
	for k, v := range counts {
		charCounts = append(charCounts, charCount{char: k, count: v})
	}

	sort.Slice(charCounts, func(i, j int) bool { return charCounts[i].count > charCounts[j].count })

	res := make([]rune, 0, len(s))
	for _, cc := range charCounts {
		for i := 0; i < cc.count; i++ {
			res = append(res, cc.char)
		}
	}

	return string(res)
}
