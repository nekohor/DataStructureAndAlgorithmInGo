package sorting

func InsertionSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	return arr
}
