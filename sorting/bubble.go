package main

func bubbleSort(arr []int) {
	length := len(arr)
	for length != 0 {
		maxIndex := 0
		for i := 1; i < length; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				maxIndex = i
			}
		}
		length = maxIndex
	}
}
