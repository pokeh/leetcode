// time complexity: O(n)
// space complexity: O(n)
func numJewelsInStones(J string, S string) int {
    jewels := make(map[rune]bool)
    for _, j := range J {
        jewels[j] = true
    }

    var cnt int
    for _, s := range S {
        if jewels[s] {
            cnt++
        }
    }
    return cnt
}
