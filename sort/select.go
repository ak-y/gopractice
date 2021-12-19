package sort

func Select(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
