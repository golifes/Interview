package LinkedList

import (
	"fmt"
	"testing"
)

func TestList_MockSingleCycleList(t *testing.T) {

	head:=s2l([]int{1,2,3,4,5})
	printNode(head)

	fmt.Println(IsCycleList(MockSingleCycleList(head)))
}
