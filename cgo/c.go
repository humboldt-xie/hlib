package main

// typedef int (intFunc) (int);
// int bridge_int_func(intFunc f,int);
import "C"

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

//export Hello
func Hello() {
}

var count int
var mtx sync.Mutex

//export Add
func Add(a, b int) int { return a + b }

//export Cosine
func Cosine(x float64) float64 { return math.Cos(x) }

//export Sort
func Sort(vals []int) { sort.Ints(vals) }

//export CallBack
func CallBack(f *C.intFunc) {
	for j := 0; j < 3; j++ {
		go func(id int) {
			for i := 0; i < 100; i++ {
				fmt.Printf("%d\n", C.bridge_int_func(f, C.int(id)))
				time.Sleep(time.Second / 10)
			}
		}(j + 1)
	}
	time.Sleep(time.Second)
}

//export Log
func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

func main() {
}
