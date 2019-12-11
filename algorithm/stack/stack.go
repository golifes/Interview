package main

import "fmt"

/**
栈:先进先出的数据结构
*/

//模拟入栈操作

func Add(y string, mid int) (s []interface{}, top int) {

	for i := 0; i <= mid; i++ {
		s = append(s, string(y[i]))
		top++
	}
	return s, top
}

func main() {

	y := "xyzyx"
	//fmt.Println(len(y))
	//for k, v := range y {
	//	fmt.Println(k, string(v))
	//}
	length := len(y)
	// bug  todo  1221 这样的
	mid := length/2 - 1
	if length%2 == 0 {
		return
	}
	add, top := Add(y, mid)

	next := mid + 2
	for i := next; i <= length-1; i++ {

		if string(y[i]) != add[top-1] {
			break
		}
		top--
	}

	if top == 0 {
		fmt.Println("回文ok")
	}
}
