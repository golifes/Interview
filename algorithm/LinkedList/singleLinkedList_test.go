package LinkedList

import (
	"fmt"
	"testing"
)

func TestList_Insert(t *testing.T) {
	list := List{}
	list.Head(1)
	list.Head(2)
	list.Head(3)
	list.Head(4)
	list.Head(5)
	list.Head(5)
	//fmt.Println("插入两个头节点-----")
	list.ShowList()
	//list.IsHui()
	//list.ShowList()
	//contains := list.Contains(3)
	//fmt.Println(contains)
	//rec := list.ContainsRec(list.headNode, 5)
	//fmt.Println(rec)
	//list.SwapPairs(list.headNode)
	elem := list.CountElem(5)
	fmt.Println(elem)

	//node := list.SwapNode(3, 4)
	//fmt.Println(node)
	//list.ShowList()

	//
	//list.Tail(3)
	//list.Tail(4)
	//
	//fmt.Println("插入两个尾节点-----")
	//list.ShowList()
	//
	//list.ShowList()
	//
	//list.RemoveNodeByValue(4)
	//fmt.Println("删除值为4的节点-----")
	//list.ShowList()
	//
	//list.RemoveNodeByValue(3)
	//fmt.Println("删除值为3的节点-----")
	//list.ShowList()
	//
	//fmt.Println("list.count", list.count)
	//
	//list.RevereList()
	//fmt.Println("链表反转-----")
	//list.ShowList()
}
