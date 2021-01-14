package main

import (
	"fmt"
	"sync"
	"time"
)

const workAmount int = 8
const timeoutLimit int = 2

func broker(ch chan<- int, id int) {

	for i := 0; i < workAmount; i++ {
		ch <- i
		fmt.Printf("Broker %v: annouced job %v\n", id, i)
		time.Sleep(time.Second)
	}
}

func worker(wg *sync.WaitGroup, ch1, ch2 <-chan int) {
	for {
		select {
		case <-time.After(time.Second * time.Duration(timeoutLimit)):
			fmt.Println("timeout")
			wg.Done()
			break
		case job := <-ch1:
			fmt.Printf("Worker: Started with job %v from ch %v\n", job, 1)
			time.Sleep(time.Second)
			fmt.Printf("Worker: Finished\n")
		case job := <-ch2:
			fmt.Printf("Worker: Started with job %v from ch %v\n", job, 2)
			time.Sleep(time.Second)
			fmt.Printf("Worker: Finished\n")
		}
	}
}

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int, workAmount)
	ch2 := make(chan int, workAmount)

	fmt.Println("Main: Starting consumer")
	wg.Add(1)
	go worker(&wg, ch1, ch2)

	fmt.Println("Main: Starting broker", 1)
	go broker(ch1, 1)
	fmt.Println("Main: Starting broker", 2)
	go broker(ch2, 2)

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}
