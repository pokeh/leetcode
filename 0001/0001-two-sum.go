package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numIdxMap := make(map[int]int)

	for i, n := range nums {
		if idx, ok := numIdxMap[target-n]; ok {
			return []int{idx, i}
		} else {
			numIdxMap[n] = i
		}
	}

	return nil
}

/* brute force
func twoSum(nums []int, target int) []int {
	for i, n1 := range nums {
		for j, n2 := range nums {
			if i != j && n1+n2 == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
*/

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 5))
}
