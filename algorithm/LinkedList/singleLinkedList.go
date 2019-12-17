package LinkedList

import "fmt"

/**
单向链表操作
*/
type SingleNode struct {
	Data interface{} //数据域
	Next *SingleNode //指向下一个节点的地址
}

//头结点
type List struct {
	headNode *SingleNode
	count    int
}

/*
	判断单向链表是否为空
*/

func (l *List) IsEmpty() bool {
	if l.count == 0 {
		return true
	}
	return false
}

//头部增加节点
func (l *List) Head(data interface{}) {
	// a b c
	// d a b c
	node := &SingleNode{Data: data}
	//把当前节点的下一个节点指向链表节点的头节点
	node.Next = l.headNode
	//把链表的头结点的头指向当前节点
	l.headNode = node

	l.count++
	return
}

//尾部增加节点
func (l *List) Tail(data interface{}) {
	node := &SingleNode{Data: data}

	if l.IsEmpty() {
		l.headNode = node

	} else {
		//获取尾节点
		tail := l.headNode

		/**
		必须判断cur.next !=nil
		因为为节点的next是nil
		*/
		for tail.Next != nil {
			tail = tail.Next
		}

		tail.Next = node //添加新的尾部节点

	}

	l.count++
	return
}

//在指定的位置添加节点
func (l *List) Insert(index int, data interface{}) {
	if index == 0 { //在头部添加节点
		l.Head(data)

	} else if index > l.count { //在尾部增加节点
		l.Tail(data)

	} else {

		cur := l.headNode //当前节点为头节点
		count := 1
		for count < index { //计算要插入位置的前置节点
			cur = cur.Next
			count += 1
		}

		/**
			a   b
		    a c b
		*/
		node := &SingleNode{Data: data}
		/*
			假如当前节点是a
			node 是c
			node.Next = cur.Next 表示c的下一个节点是b

			cur.Next = node 表示a的下一个节点是a

		*/
		node.Next = cur.Next //新增节点的后继指针指向当前节点的后继指针所指向的节点
		cur.Next = node      //当前节点的后继指针指向新增节点
	}

	l.count++
}

//删除指定位置的节点
func (l *List) RemoveNodeByIndex(index int) {

	if index == 0 { //删除头结点
		l.headNode = l.headNode.Next

	} else if index > l.count { //删除尾结点

		tail := l.headNode
		for tail.Next != nil {
			tail = tail.Next
		}

		tail.Next = nil //尾部节点next 指向nil

	} else {

		cur := l.headNode

		count := 1
		for count != index && cur.Next != nil { //寻找指定位置的节点
			count += 1
			cur = cur.Next
		}

		cur.Next = cur.Next.Next
	}

	l.count--
}

//删除指定值的节点
func (l *List) RemoveNodeByValue(data interface{}) {
	if l.count == 0 {
		return
	}

	cur := l.headNode
	if cur.Data == data { //要删除的值恰好为头节点
		// a b c  d
		//  b  c  d
		//l.headNode = cur.Next  表示把b设置为头结点
		l.headNode = cur.Next

	} else {

		for cur.Next != nil {
			if cur.Next.Data == data { //要删除的值在链表非头部位置
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
	}

	l.count--
}

//判断元素是否在节点中

func (l *List) Contains(data interface{}) bool {
	cur := l.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (l *List) ContainsRec(node *SingleNode, data interface{}) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	}

	return l.ContainsRec(node.Next, data)
}

/**
其中一个是头结点或者尾结点？ 交换数据  -->交换节点
*/
func (l *List) SwapData(x, y interface{}) bool {
	if l.count <= 1 {
		return false
	}
	if x == y {
		return false
	}

	cur := l.headNode
	curX := &SingleNode{}
	curY := &SingleNode{}

	for cur != nil {
		if cur.Data == x {
			curX = cur
		} else if cur.Data == y {
			curY = cur
		}
		cur = cur.Next
	}

	if curX.Data == nil || curY.Data == nil {
		return false
	}
	temp := curX.Data
	curX.Data = curY.Data
	curY.Data = temp
	return true
}

func (l *List) SwapNode(x, y interface{}) bool {
	if l.count <= 1 {
		return false
	}
	if x == y {
		return false
	}

	cur := l.headNode
	curX := &SingleNode{}
	curY := &SingleNode{}

	for cur != nil {
		if cur.Data == x {
			curX = cur
		} else if cur.Data == y {
			curY = cur
		}
		cur = cur.Next
	}

	if curX.Data == nil || curY.Data == nil {
		return false
	}

	//todo   bug  交换指针地址
	temp := &curX.Next
	curX.Next = curY.Next
	curY = *temp
	return true
}

func (l *List) SwapPairs(head *SingleNode) *SingleNode {
	if head == nil || head.Next == nil {
		return head
	}
	t := head.Next
	head.Next = l.SwapPairs(t.Next)
	t.Next = head
	return t
}

//遍历链表

func (l *List) ShowList() {
	if l.IsEmpty() {
		fmt.Println("link is empty")
		return
	}

	res := "head=>"

	for cur := l.headNode; cur != nil; cur = cur.Next {
		res += fmt.Sprintf("%v=>", cur.Data)
	}

	res += "tail"

	fmt.Println(res)
}

//链表反转

// 1->2->3->4->5-nil
// 5->4->3->2->1->nil
func (l *List) RevereList() {
	var prev *SingleNode
	prev = nil

	cur := l.headNode
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	l.headNode = prev
}

/**
  LeetCode
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//
//	cur := head
//	for cur != nil {
//		cur.Next, prev, cur = prev, cur, cur.Next
//	}
//	return prev
//}

func (l *List) IsHui() bool {

	var prev *SingleNode

	midLength := l.count / 2

	cur := l.headNode
	for i := 0; i < midLength; i++ {
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	midNode := cur

	var left, right *SingleNode = prev, nil
	if l.count%2 == 0 {
		right = midNode
	} else {
		right = midNode.Next
	}

	for left != nil && right != nil {
		if left.Data != right.Data {
			return false
		}
		left = left.Next
		right = right.Next
	}

	cur = prev
	prev = midNode
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	return true
}

//计算元素出现的个数
func (l *List) CountElem(x interface{}) int {

	cur := l.headNode
	if cur == nil {
		return 0
	}

	count := 0
	for cur != nil {
		if cur.Data == x {
			count++
		}
		cur = cur.Next
	}

	return count
}
