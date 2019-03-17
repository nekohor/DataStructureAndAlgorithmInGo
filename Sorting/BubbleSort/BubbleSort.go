package BubbleSort

func BubbleSort(arr []int) []int {
	length := len(arr)
	swapped := false
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			swapped = true
		}

		if !swapped {
			break
		}
	}
	return arr
}
