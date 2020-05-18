func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var window, letters [26]int

	for i := 0; i < len(s1); i++ {
		letters[s1[i]-'a']++
		window[s2[i]-'a']++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		if letters == window {
			return true
		}

		if i < len(s2)-len(s1) {
			window[s2[i]-'a']--
			window[s2[i+len(s1)]-'a']++
		}
	}

	return false
}
