package LinkedList

import (
	"fmt"
	"testing"
)

func TestList_Insert(t *testing.T) {
	list := List{}
	list.Head(1)
	list.Head(2)

	count := list.Len()
	fmt.Println(count)
}
