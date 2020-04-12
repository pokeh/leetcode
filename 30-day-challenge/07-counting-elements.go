func countElements(arr []int) int {
    cnts := make(map[int]int)

    for _, n := range arr {
        cnts[n]++
    }

    var res int

    for k, v := range cnts {
        if cnts[k+1] > 0 {
            res += v
        }
    }

    return res
}
