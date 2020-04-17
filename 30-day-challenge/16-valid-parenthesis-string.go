func checkValidString(s string) bool {
	var lo, hi int

	for _, r := range s {
		switch r {
		case '(':
			lo++
			hi++
		case ')':
			lo--
			hi--
		case '*':
			lo--
			hi++
		}

		// ran out of open brackets
		if hi < 0 {
			return false
		}

		// that * was not needed
		if lo < 0 {
			lo = 0
		}
	}

	return lo == 0
}
