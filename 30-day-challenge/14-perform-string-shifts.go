func stringShift(s string, shift [][]int) string {
	var shiftCnt int

	for _, sh := range shift {
		if sh[0] == 0 {
			shiftCnt += sh[1]
		} else {
			shiftCnt -= sh[1]
		}
	}

	for shiftCnt > len(s) {
		shiftCnt = shiftCnt - len(s)
	}

	for shiftCnt < 0 {
		shiftCnt = len(s) + shiftCnt
	}

	return s[shiftCnt:] + s[:shiftCnt]
}
