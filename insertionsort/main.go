package main

import "fmt"

func main() {
	data := [7]int{6, 4, 1, 3, 18, 24, 9}
	insertsort(data)
	// fmt.Println(data[0])
	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
}

func insertsort(data []int) {
	for i, j := 1, 0; i < len(data); i++ {
		j = i - 1
		temp := data[i]
		// if need swap
		for j >= 0 && temp < data[j] {
			data[j+1] = data[j]
			j--
		}

		//Just hold on if havn't swap
		data[j+1] = temp
	}
}
