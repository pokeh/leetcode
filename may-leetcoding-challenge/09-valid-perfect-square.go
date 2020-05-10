// O(logn) - binary search
func isPerfectSquare(num int) bool {
	lo, hi := 0, num/2+1
​
	for lo <= hi {
		mid := lo + (hi-lo)/2
​
		sq := mid * mid
		switch {
		case sq == num:
			return true
		case sq > num:
			hi = mid - 1
		case sq < num:
			lo = mid + 1
		}
	}
​
	return false
}
​
/*
// O(n)
func isPerfectSquare(num int) bool {
	i, iSq := 1, 1
​
	for {
		if iSq == num {
			return true
		}
		if iSq > num {
			return false
		}
​
		i++
		iSq = i * i
	}
}
*/
