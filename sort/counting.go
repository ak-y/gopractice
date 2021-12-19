package sort

func Counting(arr []int) []int {
	maximum := max(arr)
	counts := make([]int, maximum+1)
	ordered := make([]int, 0, len(arr))

	for _, v := range arr {
		counts[v]++
	}

	for i, v := range counts {
		for j := 0; j < v; j++ {
			ordered = append(ordered, i)
		}
	}

	return ordered
}

func max(arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	maximum := arr[0]
	for _, v := range arr {
		if v > maximum {
			maximum = v
		}
	}

	return maximum
}
