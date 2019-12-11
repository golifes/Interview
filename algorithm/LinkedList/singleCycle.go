package LinkedList

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

func MockSingleCycleList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}

	t := head

	for cur := t; cur != nil; cur = cur.Next {

		if cur.Next == nil {
			t = t.Next
			t.Next = head //最后一个节点指向头节点成环
			break
		}
	}

	return t
}

func IsCycleList(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast.Next = fast.Next.Next

		if slow == fast { //慢指针追上快指针表示有环
			return true
		}
	}

	return false
}

func s2l(s []int) *Node {
	if len(s) == 0 {
		return nil
	}

	head := &Node{
		Val: s[0],
	}

	t := head
	for i := 1; i < len(s); i++ {
		t.Next = &Node{
			Val: s[i],
		}

		t = t.Next
	}

	return head
}

func printNode(head *Node) {

	res := "head=>"

	for cur := head; cur != nil; cur = cur.Next {
		res += fmt.Sprintf("%v=>", cur.Val)
	}

	res += "tail"

	fmt.Println(res)
}

/**
  LeetCode
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	cur := head
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = cur.Next
	}

	return pre
}
