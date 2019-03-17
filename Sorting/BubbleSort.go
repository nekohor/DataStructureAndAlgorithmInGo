package sorting

// BubbleSort Implementation
func BubbleSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		swapped := false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return arr
}
