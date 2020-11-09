package test

import (
	"fmt"
	"reflect"
	"testing"
)


func judgeTypeByReflect(actual interface{}){
	var actualType = reflect.TypeOf(actual).String()

	i := 0
	if actualType != "string" && actualType != "[]interface{}" {
		// fmt.Printf("Not supported value!")
		i++
	}
}

func judgeTypeByInterface(actual interface{}){
	// switch actual.(type) {
	// case []interface{}:
	// 	fmt.Println("yes")
	// case string:
	// 	fmt.Println("yes")
	// default:
	// 	fmt.Println("no")
	// }
	_, ok1 := actual.(string)
	_, ok2 := actual.([]interface{})
	i := 0
	if !ok1 && !ok2{
		// fmt.Printf("Not supported value!")
		i++
	}
}

func BenchmarkReflect(b *testing.B){
	fmt.Printf("b.n: %v\n", b.N)
	for i := 0; i < b.N; i++ {
		judgeTypeByReflect(10)
	}
}

func BenchmarkInterface(b *testing.B){
	fmt.Printf("b.n: %v\n", b.N)
	for i := 0; i < b.N; i++ {
		judgeTypeByInterface(10)
	}
}