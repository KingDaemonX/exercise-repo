package main

import (
	section1 "exercise-repo/section_1"
	"log"
)

func main() {
	// swap
	log.Println(section1.SwapVariableByReassigning(1, 2))
	log.Println(section1.SwapVariableWithReturn(9, 7))

	// sum
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(section1.SumEvenNumberInSlice(nums))

	// interface
	file_logger := &section1.FileLogger{
		FileName: "section_1.txt",
	}

	file_logger.Log("Hello World, This is a SWE logging to a file")

	console_logger := &section1.ConsoleLogger{}
	console_logger.Log("Hello World, This is a SWE logging to the console ")
}
