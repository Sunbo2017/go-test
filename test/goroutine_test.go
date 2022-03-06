package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
// var lock sync.Mutex

type worker struct {
	name string
	age  int
}

func TestGoroutine(t *testing.T) {
	t1 := time.Now().UnixNano()
	workers := []worker{}
	for i := 0; i < 100; i++ {
		w := &worker{"w", i + 10}
		workers = append(workers, *w)
	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(w *worker, i int) {
			w.name = "w" + strconv.Itoa(i)
			fmt.Printf("over:%v\n", i)
			wg.Done()
		}(&workers[i], i)
	}
	wg.Wait()
	for _, w := range workers {
		fmt.Println(w)
	}
	t2 := time.Now().UnixNano()
	// 1000100
	fmt.Println(t2 - t1)
}

func TestNoGoroutine(t *testing.T) {
	t1 := time.Now().UnixNano()
	workers := []worker{}
	for i := 0; i < 100; i++ {
		w := &worker{"w", i + 10}
		workers = append(workers, *w)
	}
	for i := 0; i < 100; i++ {
		w := workers[i]
		w.name = "w" + strconv.Itoa(i)
		fmt.Printf("over:%v\n", i)
	}
	for _, w := range workers {
		fmt.Println(w)
	}
	t2 := time.Now().UnixNano()
	// 1000000
	fmt.Println(t2 - t1)
}

func TestGoroutineUseOneChan(t *testing.T){
	intChan := make(chan int, 12)
	// fatal error: all goroutines are asleep - deadlock!
	go goroutine1(intChan)
	go goroutine2(intChan)
	// for {
	// 	i := <- intChan
	// 	fmt.Println(i)
	// 	// if !ok {
	// 	// 	break
	// 	// }
	// }

	for i := range intChan {
		fmt.Println(i)
	}
	
	// close(intChan)
}

func goroutine1(ch chan int){
	for i:=0;i<5;i++ {
		ch <- i
	}
}

func goroutine2(ch chan int){
	for i:=0;i<5;i++ {
		ch <- i
	}
}