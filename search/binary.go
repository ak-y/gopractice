package search

func Binary(arr []int, target int) int {
	var f func(arr []int, target int, lowIndex int, hightIndex int) int
	f = func(arr []int, target int, lowIndex int, hightIndex int) int {
		if hightIndex < lowIndex {
			return -1
		}
		mid := int((lowIndex + hightIndex) / 2)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			return f(arr, target, mid+1, hightIndex)
		} else {
			return f(arr, target, lowIndex, mid-1)
		}
	}
	return f(arr, target, 0, len(arr)-1)
}

func IterBinary(arr []int, target int) int {
	lowIndex := 0
	hightIndex := len(arr) - 1
	var mid int
	for lowIndex < hightIndex {
		mid = int((lowIndex + hightIndex) / 2)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			lowIndex = mid + 1
		} else {
			hightIndex = mid
		}
	}
	return -1
}
