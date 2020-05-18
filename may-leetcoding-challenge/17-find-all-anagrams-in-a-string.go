// Sliding window
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	var res []int
	var window, letters [26]int

	for i := 0; i < len(p); i++ {
		letters[p[i]-'a']++
		window[s[i]-'a']++
	}

	for i := 0; i <= len(s)-len(p); i++ {
		if window == letters {
			res = append(res, i)
		}

		if i < len(s)-len(p) {
			window[s[i]-'a']--
			window[s[i+len(p)]-'a']++
		}
	}

	return res
}

/*
// Time limit exceeded
func findAnagrams(s string, p string) []int {
	var res []int

	p = sortStr(p)

	for i := 0; i <= len(s)-len(p); i++ {
		if sortStr(s[i:i+len(p)]) == p {
			res = append(res, i)
		}
	}

	return res
}

func sortStr(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(a, b int) bool { return runes[a] < runes[b] })
	return string(runes)
}
*/
