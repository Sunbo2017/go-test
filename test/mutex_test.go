package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var l sync.Mutex
var r sync.RWMutex

func lockDemo(s string){
	l.Lock()
	defer l.Unlock()
	for i := 0; i<5; i++{
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func TestLock(t *testing.T){
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		lockDemo("read")
	}()

	go func(){
		defer wg.Done()
		lockDemo("write")
	}()

	wg.Wait()
	fmt.Println("finish")
}

func readLock(s string){
	r.RLock()
	defer r.RUnlock()
	for i := 0; i<5; i++{
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func writeLock(s string){
	r.Lock()
	defer r.Unlock()
	for i := 0; i<5; i++{
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func TestRWLock(t *testing.T){
	var wg sync.WaitGroup
	wg.Add(4)

	go func(){
		defer wg.Done()
		fmt.Println("read1:")
		readLock("r1:")
	}()

	go func(){
		defer wg.Done()
		fmt.Println("read2:")
		readLock("r2:")
	}()

	go func(){
		defer wg.Done()
		fmt.Println("write1:")
		writeLock("w1:")
	}()

	go func(){
		defer wg.Done()
		fmt.Println("write2:")
		writeLock("w2:")
	}()

	wg.Wait()
	fmt.Println("finish")
}