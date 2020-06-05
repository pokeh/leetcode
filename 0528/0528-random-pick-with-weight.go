type Solution struct {
	sums []int
}

func Constructor(w []int) Solution {
	sums := make([]int, len(w))
	var sum int

	for i, weight := range w {
		sum += weight
		sums[i] = sum
	}

	return Solution{sums: sums}
}

func (this *Solution) PickIndex() int {
	r := rand.Intn(this.sums[len(this.sums)-1])
	return sort.Search(len(this.sums), func(i int) bool { return this.sums[i] > r })
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

/*
// Time limit exceeded
type Solution struct {
	indices []int
}

func Constructor(w []int) Solution {
	sln := Solution{}

	for i, weight := range w {
		for j := 0; j < weight; j++ {
			sln.indices = append(sln.indices, i)
		}
	}

	return sln
}

func (this *Solution) PickIndex() int {
	return this.indices[rand.Intn(len(this.indices))]
}
*/
