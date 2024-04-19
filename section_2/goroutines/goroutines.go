package main

import (
	"fmt"
	"sync"
)

func sumArray(arr []int, start int, end int, wg *sync.WaitGroup, sumChan chan int) {
	defer wg.Done()

	partialSum := 0
	for i := start; i < end; i++ {
		partialSum += arr[i]
	}
	sumChan <- partialSum
}

func main() {

	size := 100000
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i + 1
	}

	numRoutines := 4

	wg := sync.WaitGroup{}
	wg.Add(numRoutines)

	sumChan := make(chan int)

	chunkSize := size / numRoutines

	for i := 0; i < numRoutines; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == numRoutines-1 {
			end = size
		}
		go sumArray(arr, start, end, &wg, sumChan)
	}

	wg.Wait()

	totalSum := 0
	for i := 0; i < numRoutines; i++ {
		partialSum := <-sumChan
		totalSum += partialSum
	}

	fmt.Println("Total sum:", totalSum)

	close(sumChan)
}
