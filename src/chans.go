package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 10)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
			// Optional code here
		case <- time.After(time.Millisecond * 100):
			fmt.Println("Timed out.")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Println(len(c))
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}
