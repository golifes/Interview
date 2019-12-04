package slice

import "fmt"

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
