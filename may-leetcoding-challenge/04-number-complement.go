/*
// naive
func findComplement(num int) int {
	var resStr string
	for _, bit := range fmt.Sprintf("%b", num) {
		fmt.Println(bit)
		if bit == '1' {
			resStr += "0"
		} else {
			resStr += "1"
		}
	}

	res, _ := strconv.ParseInt(resStr, 2, 32)
	return int(res)
}
*/

// プロ達のonelinerを読み解く
func findComplement(num int) int {
	// eg. let num = 10010

	// 1. find the length of the binary representation of num
	// https://www.geeksforgeeks.org/count-total-bits-number/
	// => binaryLen = 5
	binaryLen := int(math.Log2(float64(num))) + 1

	// 2. create an "XOR flipper" mask of the same length as num
	// i.e. 1's with the same length as num
	// 1 << binaryLen = 100000
	// 100000 - 1 = 11111
	// => mask = 11111
	mask := 1<<binaryLen - 1

	// 3. use the XOR flipper mask to flip digits
	// 10010 ^ 11111
	// and the nature of XOR is:
	// 1 ^ 1 = 0
	// 0 ^ 1 = 1
	// => res = 01101
	res := num ^ mask

	return res
}
