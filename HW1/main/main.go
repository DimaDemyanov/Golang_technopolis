package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func swap(arr []int, i int, j int) {
	c := arr[i]
	arr[i] = arr[j]
	arr[j] = c
}

func sort(arr []int, n int) {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

func main() {
	filePathIn := "stdin"
	filePathOut := "stdout"
	fd, err := os.Open(filePathIn)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePathIn, err))
	}
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		result = append(result, x)
	}
	sort(result, len(result))
	f, err := os.Create(filePathOut) // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for i := 0; i < len(result); i++ { // Generating...
		_, err = f.WriteString(fmt.Sprintf("%d ", result[i])) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}
