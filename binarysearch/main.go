package main

import "fmt"

func main() {
	data := []int{9, 7, 8}
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i], " ")
	}
}

func binaryInsertSort(data []int) {
	var key, i, j, low, high, mid int
	for i = 1; i < len(data); i++ {
		_ = "breakpoint"
		if data[i] < data[i-1] {
			low = 0
			high = i - 1
			key = data[i]
			for low <= high {
				mid = (low + high) / 2
				if key < data[mid] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
			for j = i; j > high+1; j-- {
				data[j] = data[j-1]
			}
			data[high+1] = key

		}
	}
}
