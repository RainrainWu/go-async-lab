package main

import (
	"fmt"
	"sync"
	"time"
)

const workAmount int = 8

func broker(ch chan<- int) {

	for i := 0; i < workAmount; i++ {
		fmt.Println("Broker: annouce job", i)
		ch <- i
	}
}

func worker(wg *sync.WaitGroup, ch <-chan int, id int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started with job %v\n", id, <-ch)
	time.Sleep(time.Second)
	fmt.Printf("Worker %v: Finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, workAmount)

	for i := 0; i < workAmount; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go worker(&wg, ch, i)
	}

	fmt.Println("Main: Starting broker")
	go broker(ch)

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}
