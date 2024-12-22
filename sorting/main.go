package main

import "fmt"

func main() {
	arr := [...]int{3, 6, 7, 3, 2, 1} //   3 6 7 2 1    1 6 7 2 3   1 3 7 2 6   1 3 6 2 7
	//   3 6 7 2 1
	// 3
	//bubbleSort(arr[:])
	selectionSort(arr[:])
	fmt.Println(arr)
}
