func reverseWords(s string) string {
	var b strings.Builder

    arr := strings.Split(s, " ")

    for i := len(arr)-1; i >= 0; i-- {
        if arr[i] != "" {
			b.WriteString(arr[i])
			b.WriteString(" ")
		}
    }

	return strings.TrimSpace(b.String())
}
