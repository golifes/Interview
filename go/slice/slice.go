package main

import "fmt"

func main() { //expand cap
	a := make([]int, 1024) //2,1024,5000
	fmt.Printf("len(%d),cap(%d)\n", len(a), cap(a))

	//a = append(a, 1,2,3)

	a = append(a, 1)
	fmt.Printf("len(%d),cap(%d)\n", len(a), cap(a))

	a = append(a, 2, 3, 4)
	fmt.Printf("len(%d),cap(%d)\n", len(a), cap(a))
	a = append(a, 5, 6)
	fmt.Printf("len(%d),cap(%d)\n", len(a), cap(a))
	//
	a = append(a, 7, 8, 9, 10)
	fmt.Printf("len(%d),cap(%d)\n", len(a), cap(a))
	//a := 1
	//c := a + 2
	//a = 4
	//fmt.Println(c)

	/**

	 	growslice(et *_type, old slice, cap int)
			cap:表示当前slice的长度+append的元素个数
		eg:
			a := make([]int, 1)
			a=append(a, 1)
			growslice(et *_type, old slice, cap int)就可以表示如下:
				cap=2
				old.cap = 1
				newcap := old.cap = 1
				doublecap := newcap + newcap = 2
					1      2
				if cap > doublecap {
					newcap = cap
				} else {
					if old.len < 1024 {
						newcap = doublecap
						newcap = 2
			 ....下面进行lenmem和capmem 以及新的newcap的操作计算
			此时newcap是2，len是2
		   a=append(a, 1,2,3)
			growslice(et *_type, old slice, cap int)就可以表示如下:
				cap= 2 + 3  =5
				old.cap = 2
				newcap := old.cap = 2
				doublecap := newcap + newcap = 4
					5      4
				if cap > doublecap {
					newcap = cap = 5
				} else {
					if old.len < 1024 {
						newcap = doublecap
			此时:newcap就是5,len=newcap

	*/
}

//func Create() {
//	s1 := []string{"a", "b", "c", "d"}
//	s2 := make([]string, 5, 5)
//	s3 := make([]string, 4)
//	fmt.Println(s1)
//	//fmt.Println(s2, len(s2), cap(s2))
//	fmt.Printf("%#v　 ,%d,%d", s2, len(s2), cap(s2))
//	fmt.Printf("%#v　 ,%d,%d", s3, len(s3), cap(s3))
//	fmt.Println(s3[4:])
//
//}
//
////改变slice  array也会改变
//func Array() {
//	var array = [...]int64{1, 2, 3, 4, 5}
//	var slice = array[2:4]
//
//	fmt.Printf("改变前---> %+v  ,%+v", array, slice)
//	slice[0] = 100
//	fmt.Printf("改变后---> %+v  ,%+v", array, slice)
//
//}
//
//func Addr() {
//	s := []int{1, 2}
//	for i := 0; i < 17; i++ {
//		fmt.Printf("addr %p,len %d ,cap %d\n", &s, len(s), cap(s))
//		s = append(s, i)
//	}
//}
