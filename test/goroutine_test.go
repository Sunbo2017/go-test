package test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

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
