func backspaceCompare(S string, T string) bool {
	return text(S) == text(T)
}

func text(keys string) string {
	var res string
	for _, r := range keys {
		if r == '#' {
			if len(res) != 0 {
				res = res[:len(res)-1]
			}
		} else {
			res += string(r)
		}
	}
	return res
}
