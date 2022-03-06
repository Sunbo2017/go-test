package test

import (
	"testing"
	"fmt"
)

type position struct {
	x int
	y int
}

var total int32

func queen(size int) {
    boards := make([]int, size)
    //第0列开始放queen
    put(boards, 0)
    fmt.Println(total)
}

func put(boards []int, col int) {
    size := len(boards)
    if col == size {
        fmt.Println(boards) // 输出答案
        total++
        return
    }

    for row := 0; row < size; row++ {
		// 在 row 处放下皇后,即坐标为（col,row)
        boards[col] = row 
        if safe(boards, col) {
            put(boards, col+1)
        }
    }
}

func safe(boards []int, col int) bool {
    for c := 0; c < col; c++ {
        if isAttack(boards, c, col) {
            return false
        }
    }
    return true
}

func isAttack(boards []int, c int, col int) bool {
    switch {
    case boards[c] == boards[col]://在同一行
        return true
    case boards[col]-boards[c] == c-col://下斜线
        return true
    case boards[col]-boards[c] == col-c://上斜线
        return true
    }
    return false
}

func TestQueen(t *testing.T) {
	queen(8)
}