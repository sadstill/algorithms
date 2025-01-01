package main

func quickSort(arr []int) {
	startQuickSort(arr, 0, len(arr)-1)
}

func startQuickSort(arr []int, lowIndex, highIndex int) {
	if lowIndex >= highIndex {
		return
	}

	pivot := arr[highIndex]

	l := partition(arr, lowIndex, highIndex, pivot)

	startQuickSort(arr, lowIndex, l-1)
	startQuickSort(arr, l+1, highIndex)
}

func partition(arr []int, lowIndex, highIndex, pivot int) int {
	l := lowIndex
	r := highIndex - 1

	for l < r {
		for arr[l] <= pivot && l < r {
			l++
		}
		for arr[r] >= pivot && l < r {
			r--
		}
		arr[l], arr[r] = arr[r], arr[l]
	}

	if arr[l] > arr[highIndex] {
		arr[l], arr[highIndex] = arr[highIndex], arr[l]
	} else {
		l = highIndex
	}

	return l
}

// 5, 3, 8, 4, 2
// 2, 3, 8, 4, 5
// 3, 8, 4, 5
// 3, 4, 5, 8
// 3, 4, 5
// 3, 4, 5
// 3, 4
// 3

// 3 4
