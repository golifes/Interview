package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	var b interface{} = (*interface{})(nil)
	var c = (interface{})(nil)

	fmt.Println(a, a == nil)
	fmt.Println(b, b == nil)
	fmt.Println(c, c == nil)

	fmt.Printf("a type: %#v\t value: %#v \n",
		reflect.TypeOf(a),
		reflect.ValueOf(a).String(),
	)

	fmt.Printf("b type: %#v\t value: %#v \n",
		reflect.TypeOf(b).String(),
		reflect.ValueOf(b).String(), //*(*[2]uintptr)(unsafe.Pointer(&b))
	)

	fmt.Printf("c type: %#v\t value: %#v \n",
		//reflect.TypeOf(c).String(),
		reflect.ValueOf(c).String(), //*(*[2]uintptr)(unsafe.Pointer(&b))
	)
}
