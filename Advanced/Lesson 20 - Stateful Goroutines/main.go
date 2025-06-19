package main

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	Id    int
	count int
	ch    chan int
}

func (w *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println(w.Id, "Current count:", w.count)
			}
		}
	}()
}

func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func (w *StatefulWorker) Close() {
	close(w.ch)
}

func main() {

	stWorker := &StatefulWorker{
		Id: 1,
		ch: make(chan int),
	}

	stWorker2 := &StatefulWorker{
		Id: 2,
		ch: make(chan int),
	}

	stWorker.Start()
	stWorker2.Start()

	for i := range 5 {
		stWorker.Send(i + 1)
		stWorker2.Send((i + 1) * 2)
		time.Sleep(500 * time.Millisecond)
	}

	stWorker.Close()
	stWorker2.Close()

}
