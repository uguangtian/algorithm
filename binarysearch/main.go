package main

import "fmt"

func main() {
	fmt.Println("Runing")
	data := []int{2, 8, 22, 144, 222}
	index := binarySearch(data, 222)
	fmt.Println(index)
}

func binarySearch(data []int, goal int) int {
	low := 0
	high := len(data) - 1
	for low <= high {
		middle := low + ((high - low) >> 1)
		if goal == data[middle] {
			return middle
		} else if goal < data[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
		return -1
	}
	return 0
}
