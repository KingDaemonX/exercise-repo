package main

import (
	section1 "exercise-repo/section_1"
	"log"
)

func main() {
	log.Println(section1.SwapVariableByReassigning(1, 2))
	log.Println(section1.SwapVariableWithReturn(9, 7))

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(section1.SumEvenNumberInSlice(nums))
}
