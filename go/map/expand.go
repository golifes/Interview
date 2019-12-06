package main

import "fmt"

var  (
	// makehmap_small implements Go map creation for make(map[k]v) and
	// make(map[k]v, hint) when hint is known to be at most bucketCnt
	// at compile time and the map needs to be allocated on the heap.
	b = make(map[int]int)//在堆上分配内存，调用 makemap_small
)

func main() {
	var a map[int]int

	// Only need to initialize h.hash0 since
	// hmap h has been allocated on the stack already.
	// h.hash0 = fastrand()
	a = make(map[int]int) //创建map，在栈上分配直接调用 fastrand

	a[1] = 1      //map 存值 runtime.mapassign_fast64(SB)
	v, ok := a[1] //map取值

	fmt.Println(v, ok)
}

