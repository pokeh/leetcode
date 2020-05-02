/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// time complexity: O(logn)
// space complexity: O(1)
func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		// use this instead of (left + right) / 2 to prevent overflow
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
