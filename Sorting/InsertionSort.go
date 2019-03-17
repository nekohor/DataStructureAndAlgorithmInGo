package sorting

func InsertionSort(arr []int) []int {
	length := len(arr)
	var min int
	for i := 0; i < length; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
	return arr
}
