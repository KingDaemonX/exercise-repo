package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		number := rand.Intn(100)
		ch <- number
		fmt.Println("Produced:", number)
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch chan int) {
	for num := range ch {
		square := num * num
		fmt.Println("Consumed:", square)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	go producer(ch)

	go consumer(ch)

	time.Sleep(time.Second * 33)
}
