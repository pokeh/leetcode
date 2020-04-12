func isHappy(n int) bool {
    nums := make(map[int]struct{})

    for {
        n = sumOfSquares(n)

        if n == 1 {
            return true
        }

        if _, ok := nums[n]; ok {
            return false
        }

        nums[n] = struct{}{}
    }
}

func sumOfSquares(n int) int {
    sum := 0

    for n != 0 {
        sum += (n%10) * (n%10)
        n /= 10
    }

    return sum
}
