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
}

/*
	判断单向链表是否为空
*/

func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	}
	return false
}

//获取长度
func (l *List) Len() (count int) {
	cur := l.headNode
	for cur != nil {
		count += 1
		cur = cur.Next
	}
	return count
}

/**
添加元素(尾部添加和头部添加)
*/

func (l *List) Head(data interface{}) *SingleNode {
	// a b c
	// d a b c
	node := &SingleNode{Data: data}
	//把当前节点的下一个节点只想链表节点的头节点
	node.Next = l.headNode
	//把链表的头结点的头指向当前节点
	l.headNode = node
	return node
}

func (l *List) Tail(data interface{}) *SingleNode {
	node := &SingleNode{Data: data}
	if l.IsEmpty() {
		l.headNode = node
	} else {
		//获取尾节点
		cur := l.headNode

		/**
		必须判断cur.next !=nil
		*/
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
	return node
}

//在指定的位置添加元素
func (l *List) Insert(index int, data interface{}) *SingleNode {
	if index <= 0 {
		return l.Head(data)
	} else if index > l.Len() {
		return l.Tail(data)
	} else {
		//前置节点
		cur := l.headNode
		count := 1
		for count < index {
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
		return node
	}
}

//删除指定位置的元素
func (l *List) RemoveIndex(index int) {
	//保证这个位置是有效的
	if index == 0 {
		//删除头结点
		l.headNode = l.headNode.Next
	} else if index < 0 || index > l.Len() {
		return
	} else {
		count := 1
		cur := l.headNode
		for count != index && cur.Next != nil {
			count += 1
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
}

//删除指定的值
func (l *List) Remove(data interface{}) {
	cur := l.headNode
	if cur.Data == data {
		// a b c  d
		//  b  c  d
		//l.headNode = cur.Next  表示把b设置为头结点
		l.headNode = cur.Next
	} else {
		//循环遍历
		for cur.Next != nil {
			if cur.Next.Data == data {
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
	}
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
	if l.Len() != 0 {
		cur := l.headNode

		for {
			fmt.Sprintf("\t%v", cur.Data)
			if cur != nil {
				cur = cur.Next
			}
		}
	}
}
