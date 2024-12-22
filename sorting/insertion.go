package main

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		sorted := i - 1
		for sorted > -1 && arr[sorted] > arr[sorted+1] {
			arr[sorted], arr[sorted+1] = arr[sorted+1], arr[sorted]
			sorted--
		}
	}
}
