package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	msgs := make(chan int, 100)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			msgs <- i
		}
	}()
	wg.Wait()

	collects := make(chan int, 5)
	go func() {
		for {
			msg := <-collects
			fmt.Println("collect msg=", msg)
			time.Sleep(time.Second)
		}
	}()

	var run = true
	for run {
		select {
		case msg := <-msgs:
			fmt.Println("i=", msg)
			collects <- msg
		default:
			fmt.Println("exit select")
			run = false
		}
	}

	fmt.Println("end")
	time.Sleep(10000 * time.Second)
}
