func firstUniqChar(s string) int {
	runeCnt := [26]int{}
​
	for _, r := range s {
		runeCnt[r-'a']++
	}
​
	for i, r := range s {
		if runeCnt[r-'a'] == 1 {
			return i
		}
	}
	return -1
}
​
/*
func firstUniqChar(s string) int {
	// note: index starts from 1, to distinguish from the initial value
	uniqRuneIdx := [26]int{}
​
	for i, r := range s {
		switch uniqRuneIdx[r-'a'] {
		case -1:
			continue
		case 0:
			uniqRuneIdx[r-'a'] = i + 1
		default:
			uniqRuneIdx[r-'a'] = -1
		}
	}
​
	maxInt := int(^uint(0) >> 1)
	firstIdx := maxInt
​
	for _, idx := range uniqRuneIdx {
		if idx > 0 && idx < firstIdx {
			firstIdx = idx
		}
	}
​
	if firstIdx == maxInt {
		return -1
	}
	return firstIdx - 1 // convert back to 0-start index
}
*/
