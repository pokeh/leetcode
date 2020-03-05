package main

import "fmt"

func readBinaryWatch(num int) []string {
	res := make([]string, 0)

	for hNum := 0; hNum <= num; hNum++ {
		mNum := num - hNum

		// fmt.Printf("listPossibleHrs with hNum=%v\n", hNum)
		hList := listPossibleHrs(hNum)
		// fmt.Printf("listPossibleMins with mNum=%v\n", mNum)
		mList := listPossibleMins(mNum)

		res = append(res, listTimeCombinations(hList, mList)...)
	}

	return res
}

func listPossibleHrs(hNum int) []int {
	res := [][]int{}

	hours := []int{1, 2, 4, 8}
	bound := 11

	listCombinations(hNum, hours, bound, &res)
	return sumSliceSlice(res)
}

func listPossibleMins(mNum int) []int {
	res := [][]int{}

	mins := []int{1, 2, 4, 8, 16, 32}
	bound := 59

	listCombinations(mNum, mins, bound, &res)
	return sumSliceSlice(res)
}

func sumSliceSlice(ss [][]int) []int {
	res := []int{}

	for _, s := range ss {
		res = append(res, sumSlice(s))
	}

	return res
}

func sumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

func listCombinations(k int, ns []int, bound int, res *[][]int) {
	helper(0, k, ns, bound, []int{}, res)
}

func helper(pointer, k int, ns []int, bound int, curr []int, res *[][]int) {
	// fmt.Printf("pointer=%v, k=%v, curr=%v, res=%v\n", pointer, k, curr, res)

	if len(curr) == k {
		*res = append(*res, append([]int{}, curr...))
		return
	}

	if pointer == len(ns) {
		return
	}

	if sumSlice(curr)+ns[pointer] > bound {
		return
	}

	for i := pointer; i < len(ns); i++ {
		helper(i+1, k, ns, bound, append(curr, ns[i]), res)
	}
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
	// hList := []int{1, 2, 4}
	// mList := []int{3, 5, 9, 17, 33}
	// fmt.Println(listTimeCombinations(hList, mList))

	// fmt.Println(listPossibleHrs(2))  // should exclude 12
	// fmt.Println(listPossibleMins(4)) // should exclude 60

	fmt.Println(readBinaryWatch(2))
}
