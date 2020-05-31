func kClosest(points [][]int, K int) [][]int {
	pointsByDistance := make(map[float64][][]int)

	for _, p := range points {
		d := distance(p)
		pointsByDistance[d] = append(pointsByDistance[d], p)
	}

	keys := make([]float64, 0, len(pointsByDistance))
	for key := range pointsByDistance {
		keys = append(keys, key)
	}
	sort.Float64s(keys)

	res := make([][]int, 0, K)

	for i, j := 0, 0; i < K; {
		v, _ := pointsByDistance[keys[j]]
		res = append(res, v...)
		j++
		i += len(v)
	}

	return res
}

func distance(point []int) float64 {
	return math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
}
