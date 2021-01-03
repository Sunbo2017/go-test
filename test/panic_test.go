package test

import (
	"fmt"
	"testing"
	"time"
)

// 一个协程发生panic会影响其他协程，且无法在其他协程捕获异常
// 只能由发生异常的协程自己捕获处理异常
func TestPanic(t *testing.T){
	fmt.Println("start")

	go func ()  {
		for i := 0;i<10;i++{
			fmt.Println("run goroutine2")
		}
	}()

	go func ()  {
		fmt.Println("run goroutine1")
		panic("goroutine panic")
	}()

	time.Sleep(time.Second)

	fmt.Println("end")
}