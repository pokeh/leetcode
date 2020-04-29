type FirstUnique struct {
	maybeUniqs []int
	counts     map[int]int
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{
		maybeUniqs: make([]int, 0),
		counts:     make(map[int]int),
	}
	for _, n := range nums {
		fu.Add(n)
	}
	return fu
}

func (this *FirstUnique) ShowFirstUnique() int {
	for i, n := range this.maybeUniqs {
		if this.counts[n] == 1 {
			this.maybeUniqs = this.maybeUniqs[i:]
			return this.maybeUniqs[0]
		}
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	if this.counts[value] == 0 {
		this.maybeUniqs = append(this.maybeUniqs, value)
	}
	this.counts[value]++
}

/*
type FirstUnique struct {
	uniqueNums []int
	numCounts  map[int]int
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{
		uniqueNums: make([]int, 0),
		numCounts:  make(map[int]int),
	}

	for _, n := range nums {
		fu.Add(n)
	}

	return fu
}

func (this *FirstUnique) ShowFirstUnique() int {
	if len(this.uniqueNums) == 0 {
		return -1
	}
	return this.uniqueNums[0]
}

func (this *FirstUnique) Add(value int) {
	currCnt := this.numCounts[value]
	if currCnt == 0 {
		this.markAsUnique(value)
	} else {
		this.markNonUnique(value)
	}
	this.numCounts[value]++
}

func (this *FirstUnique) markAsUnique(value int) {
	this.uniqueNums = append(this.uniqueNums, value)
}

func (this *FirstUnique) markNonUnique(value int) {
	for i, n := range this.uniqueNums {
		if n == value {
			this.uniqueNums = append(this.uniqueNums[:i], this.uniqueNums[i+1:]...)
		}
	}
}
*/
