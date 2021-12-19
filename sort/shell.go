package sort

func Shell(arr []int) []int {
	len := len(arr)
	for gap := len / 2; gap > 0; gap /= 2 {
		for i := 0; i < len-gap; i++ {
			for j := i; j >= 0 && arr[j] > arr[j+gap]; j -= gap {
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
	}
	return arr
}
