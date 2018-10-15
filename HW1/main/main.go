package main

import "fmt"

func swap(arr []float64, i int, j int) {
	c := arr[i]
	arr[i] = arr[j]
	arr[j] = c
}

func sort(arr []float64, n int) {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

func main() {
	length := 0
	fmt.Println("Enter the number of inputs")
	fmt.Scan(&length)
	fmt.Println("Enter the inputs")
	numbers := make([]float64, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&numbers[i])
	}
	sort(numbers, length)
	fmt.Println(numbers)
}
