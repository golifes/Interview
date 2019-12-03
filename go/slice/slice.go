package slice

import "fmt"

func Create() {
	s1 := []string{"a", "b"}
	s2 := make([]string, 5, 5)
	s3 := make([]string, 4)
	fmt.Println(s1)
	//fmt.Println(s2, len(s2), cap(s2))
	fmt.Printf("%#v　 ,%d,%d", s2, len(s2), cap(s2))
	fmt.Printf("%#v　 ,%d,%d", s3, len(s3), cap(s3))
	fmt.Println(s3[4:])

}
