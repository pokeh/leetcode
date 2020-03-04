package main

import "fmt"

func readBinaryWatch(num int) []string {
	res := make([]string, 0)

	for hNum := 0; hNum <= num; hNum++ {
		mNum := num - hNum

		hList := listPossibleHrs(hNum)
		mList := listPossibleMins(mNum)

		res = append(res, listTimeCombinations(hList, mList)...)
	}

	return res
}

func listPossibleHrs(hNum int) []int {
	res := []int{}

	hours := []int{1, 2, 4, 8}
	max := 11

	helper(0, hNum, hours, max, &res)
	return res
}

func listPossibleMins(mNum int) []int {
	res := []int{}

	mins := []int{1, 2, 4, 8, 16, 32}
	max := 59

	helper(0, mNum, mins, max, &res)
	return res
}

func helper(currSize int, lastSize int, list []int, max int, res *[]int) {
	if currSize == lastSize {
		return
	}

	// TODO
}

func listTimeCombinations(hList, mList []int) []string {
	res := make([]string, 0, len(hList)*len(mList))

	for _, h := range hList {
		for _, m := range mList {
			res = append(res, fmt.Sprintf("%d:%02d", h, m))
		}
	}

	return res
}

func main() {
	hList := []int{1, 2, 4}
	mList := []int{3, 5, 9, 17, 33}
	fmt.Println(listTimeCombinations(hList, mList))

	// fmt.Println(listPossibleHrs(2)) // should exclude 12
	// fmt.Println(listPossibleMins(4)) // should exclude 60
}
