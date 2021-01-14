package main

import (
	"fmt"
	"sync"
	"time"
)

const workAmount int = 4

func broker(ch chan<- int, amount, id int) {

	for i := 0; i < amount; i++ {
		ch <- i
		fmt.Printf("Broker %v: annouced job %v\n", id, i)
	}
}

func worker(wg *sync.WaitGroup, ch1 <-chan int, ch2 <-chan int, id int) {
	defer wg.Done()

	select {
	case job := <-ch1:
		fmt.Printf("Worker %v: Started with job %v from ch %v\n", id, job, 1)
		time.Sleep(time.Second)
		fmt.Printf("Worker %v: Finished\n", id)
	case job := <-ch2:
		fmt.Printf("Worker %v: Started with job %v from ch %v\n", id, job, 2)
		time.Sleep(time.Second)
		fmt.Printf("Worker %v: Finished\n", id)
	}
}

func main() {
	var wg sync.WaitGroup

	halfAmount := int(workAmount / 2)
	ch1 := make(chan int, halfAmount)
	ch2 := make(chan int, workAmount-halfAmount)

	for i := 0; i < workAmount; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go worker(&wg, ch1, ch2, i)
	}

	fmt.Println("Main: Starting broker", 1)
	go broker(ch1, halfAmount, 1)
	fmt.Println("Main: Starting broker", 2)
	go broker(ch2, workAmount-halfAmount, 2)

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}
