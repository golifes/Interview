package main

import "fmt"

func main() { //expand cap
	a := make([]int, 2)//2,1024,5000
	fmt.Printf("len(%d),cap(%d)\n", len(a), cap(a))

	//a = append(a, 1,2,3)

	//扩容逻辑：
	// cap=old.cap+1= 2+1 =3
	// newcap := old.cap = 2
	// doublecap :=  newcap + newcap = 2+2 =4

	// 如果 cap > doublecap 则 newcap = cap = 3,
	// 否则 cap < doublecap， 如果old.len<1024 则 newcap = doublecap =4

	//内存对齐 roundupsize
	//roundupsize(uintptr(newcap) * sys.PtrSize) = roundupsize(4*8) =roundupsize(32)
	// 32<_MaxSmallSize=32768， 并且 size <= smallSizeMax-8 = 32<=1024-8=1016
	// 走逻辑 return uintptr(class_to_size[size_to_class8[(size+smallSizeDiv-1)/smallSizeDiv]])

	// size_to_class8[(size+smallSizeDiv-1)/smallSizeDiv] = size_to_class8[(32+8-1)/8] = size_to_class8[(32+8-1)/8] = size_to_class8[4]
	// 查size_to_class8 数组得 size_to_class8[4]= 3， 再查class_to_size数组 得class_to_size[3] = 32

	// 计算 capmem = roundupsize(uintptr(newcap) * sys.PtrSize) = roundupsize(32) =32
	// 最后计算 newcap = int(capmem / sys.PtrSize) = int(32/8)=4

	//最终返回 return slice{p, old.len, newcap} => return slice{p, old.len, 4}

	a=append(a, 1)//所以cap=4，同理可计算下面的cap
	a=append(a, 2)
	a=append(a, 3)

	fmt.Printf("len(%d),cap(%d)\n", len(a), cap(a))
}

func Create() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := make([]string, 5, 5)
	s3 := make([]string, 4)
	fmt.Println(s1)
	//fmt.Println(s2, len(s2), cap(s2))
	fmt.Printf("%#v　 ,%d,%d", s2, len(s2), cap(s2))
	fmt.Printf("%#v　 ,%d,%d", s3, len(s3), cap(s3))
	fmt.Println(s3[4:])

}

//改变slice  array也会改变
func Array() {
	var array = [...]int64{1, 2, 3, 4, 5}
	var slice = array[2:4]

	fmt.Printf("改变前---> %+v  ,%+v", array, slice)
	slice[0] = 100
	fmt.Printf("改变后---> %+v  ,%+v", array, slice)

}

func Addr() {
	s := []int{1, 2}
	for i := 0; i < 17; i++ {
		fmt.Printf("addr %p,len %d ,cap %d\n", &s, len(s), cap(s))
		s = append(s, i)
	}
}
