package sort

func Comb(arr []int) []int {
	var (
		l       = len(arr)
		gap     = int(float64(l) / 1.3)
		swapped = true
	)

	for gap != 1 || swapped {
		swapped = false

		for i := 0; i < l-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
		}

		gap = int(float64(gap) / 1.3)
		if gap < 1 {
			gap = 1
		}
	}

	return arr
}
