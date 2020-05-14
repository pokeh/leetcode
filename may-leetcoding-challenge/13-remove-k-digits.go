// thanks to https://www.youtube.com/watch?v=HXvxIth510g
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	stack := []byte{num[0]}
	remainingRemove := k

	// use a stack to select the smallest elements for the top digits
	for i := 1; i < len(num); i++ {
		// while we are allowed to remove elements,
		// pop any preceding elements that are larger than the current
		for remainingRemove > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			remainingRemove--
		}
		stack = append(stack, num[i])
	}

	wantLen := len(num) - k
	if len(stack) > wantLen {
		stack = stack[0:wantLen]
	}

	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		return "0"
	}
	return res
}
