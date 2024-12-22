package main

import "fmt"

func main() {
	arr := [...]int{3, 6, 7, 3, 2, 1}
	//bubbleSort(arr[:])
	insertionSort(arr[:])
	//selectionSort(arr[:])
	fmt.Println(arr)
}
