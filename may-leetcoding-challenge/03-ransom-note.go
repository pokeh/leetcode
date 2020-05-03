// time complexity: O(n)
// space complexity: O(n)
func canConstruct(ransomNote string, magazine string) bool {
	// optimization
	if len(ransomNote) > len(magazine) {
		return false
	}

	availableRunes := make(map[rune]int)
	for _, r := range magazine {
		availableRunes[r]++
	}

	for _, r := range ransomNote {
		if availableRunes[r] == 0 {
			return false
		}
		availableRunes[r]--
	}

	return true
}
