package main

import "fmt"

func main() {
	data := []int{6, 4, 1, 3, 18, 24, 9}
	insertsort(data)

	for i := 1; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
}

func insertsort(data []int) {
	n := data[0]
	for i, j := 2, 0; i < n+1; i++ {
		j = i - 1
		temp := data[i]
		for j > 0 && temp < data[j] {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}
