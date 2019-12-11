package LinkedList

import (
	"fmt"
	"testing"
)

func TestList_Insert(t *testing.T) {
	list := List{}
	list.Head(1)
	list.Head(2)

	fmt.Println("插入两个头节点-----")
	list.ShowList()

	list.Tail(3)
	list.Tail(4)

	fmt.Println("插入两个尾节点-----")
	list.ShowList()

	list.ShowList()

	list.RemoveNodeByValue(4)
	fmt.Println("删除值为4的节点-----")
	list.ShowList()

	list.RemoveNodeByValue(3)
	fmt.Println("删除值为3的节点-----")
	list.ShowList()

	fmt.Println("list.count", list.count)

	list.RevereList()
	fmt.Println("链表反转-----")
	list.ShowList()
}
