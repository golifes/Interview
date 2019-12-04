package slice

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	Create()
}

func TestArray(t *testing.T) {
	Array()
}

func TestAddr(t *testing.T) {
	Addr()
	u := ^uintptr(0)
	fmt.Println(u)
	//4 << (^uintptr(0) >> 63)
	//按位异或 0  = 2^64次方-1
	a := 4 << (^uintptr(0) >> 63)
	b := 18446744073709551615 >> 63
	c := 4 << 1
	d := 1 << 8
	// 0 0 0 0 0 0 0 0
	fmt.Println(1 | 2)
	// 0 1
	// 1 0
	// 1 1
	fmt.Println(a, b, c, d)

}
