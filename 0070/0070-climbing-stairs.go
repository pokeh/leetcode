package main

// Question

// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
// Note: Given n will be a positive integer.

// Thought process

// what is the generalized question?
// => how many distinct ways can you climb to the k+1-th stair?

// # 1:
// how many distinct ways can you climb to the 1st stair?
// ans: 1
// 1 step

// # 2:
// how many distinct ways can you climb to the 2nd stair?
// ans: 2
// 1 step + 1 step
// 2 step

// # 3:
// how many distinct ways can you climb to the 3rd stair?
// ans: 3
// 1 step + 2 step
// 1 step + 1 step + 1 step
// 2 step + 1 step

// # 4:
// how many distinct ways can you climb to the 4th stair?
// ans: 5
// 1 step + 1 step + 1 step + 1 step
// 1 step + 2 step + 1 step
// 2 step + 1 step + 1 step
// 1 step + 1 step + 2 step
// 2 step + 2 step

// Can you find the number of distinct ways of k, given answers of 1...k-1?
// Yes!
// The total number of distinct ways are ans(k-1) + ans(k-2).
// This is because, the ways to reach the k-th stair,
// we could add 1 step to each answer for k-1,
// or, we could add 2 steps to each answer for k-2. Those are the only answers.

// Now that we have this, we could just solve this problem recursively.

var memo []int

func climbStairs(n int) int {
	memo = make([]int, n+1)

	return climb(n)
}

func climb(k int) int {
	if k < 0 {
		panic("unexpected")
	}

	if k < 2 {
		memo[k] = 1
		return 1
	}

	if memo[k] != 0 {
		return memo[k]
	}

	ans := climb(k-1) + climb(k-2)
	memo[k] = ans
	return ans
}
