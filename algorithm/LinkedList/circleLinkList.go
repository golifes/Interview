package LinkedList

//节点部分
type CirCleNode struct {
	Data interface{}
	Pre  *CirCleNode
	Next *CirCleNode
}

//节点
type CirCleList struct {
	Count int
	Head  *CirCleNode //头节点
	Tail  *CirCleNode //尾结点
}

//初始化工作
func (c *CirCleList) New() {
	c.Count = 0
	c.Head = nil
	c.Tail = nil
}

//查询下标
func (c *CirCleList) Index(index int) *CirCleNode {
	if c.Count == 0 || index > c.Count-1 {
		return nil
	}

	if index == 0 {
		return c.Head
	}
	//遍历
	node := c.Head

	for i := 0; i <= index; i++ {
		node = node.Next
	}

	return node
}

//后面添加
func (c *CirCleList) Append(node *CirCleNode) bool {
	if node == nil {
		return false
	}
	//就一个节点
	if c.Count == 0 {
		c.Head = node
		c.Tail = node
		node.Next = nil
		node.Pre = nil
	} else {

		/**
				     cur      node
		 			 a        b

					在a后面添加b
					b的前置节点是a尾结点，也就是a的尾结点指向的是b的前置节点
					b的下一个节点

		*/
		node.Pre = c.Tail
		//node.Next = nil
		c.Tail.Next = node
		c.Tail = node
	}

	c.Count++

	return true
}

//指定位置插入
func (c *CirCleList) Insert(index int, node *CirCleNode) bool {
	//-1 表示尾部 大于count就是超过了尾部
	if node == nil || index > c.Count {
		return false
	}

	if index == c.Count {
		return c.Append(node)
	}

	if index == 0 {
		node.Next = c.Head
		c.Head = node
		//c.Head.Pre = nil
		c.Count++
		return true
	}

	cleNode := c.Index(index)
	node.Pre = cleNode.Pre
	node.Next = cleNode
	cleNode.Pre.Next = node
	cleNode.Pre = node
	c.Count++
	return true
}

//删除节点
func (c *CirCleList) Delete(index int) bool {
	if index > c.Count-1 {
		return false
	}

	if index == 0 {
		if c.Count == 1 {
			c.Head = nil
			c.Tail = nil
		} else {
			// todo ????
			c.Head.Next.Pre = nil
			c.Head = c.Head.Next
		}

		c.Count--
		return true
	}

	node := c.Index(index)
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	c.Count--
	return true
}
