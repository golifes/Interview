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
		node.Next = cur.Next
		cur.Next = node
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
