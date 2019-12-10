package LinkedList

//节点部分
type CirCleNode struct {
	Data interface{}
	Pre  *CirCleNode
	Next *CirCleNode
}

//节点 http://www.360doc.com/content/18/0518/15/11935121_754982751.shtml
type CirCleList struct {
	Count int
	Head  *CirCleNode //头节点
	Tail  *CirCleNode //尾结点
}

type CirClear interface {
	New()
	GetIndex(index int) *CirCleNode          //查询下表
	Append(node *CirCleNode) bool            //尾部添加
	Insert(index int, node *CirCleNode) bool //指定位置插入节点数据
	Delete(index int) bool                   //按照下标删除
	IsHead(node *CirCleNode) bool            //判断是否是头结点
	IsTail(node *CirCleNode) bool            //判断是否是头结点

}

type CirNodes interface {
	GetNodeData() interface{} //获取某个节点的数据
	GetNext() *CirCleNode     //获取当前节点的下一个节点
	GetPre() *CirCleNode      //获取当前节点的上一个节点
}

func (c *CirCleList) New() {
	c.Count = 0
	c.Head = nil //头
	c.Tail = nil //尾
}

func (c *CirCleList) GetIndex(index int) *CirCleNode {
	if c.Count == 0 || index > c.Count-1 {
		return nil
	}

	if index == 0 {
		return c.Head
	}

	node := c.Head
	count := 0
	for count < index {
		node = node.Next
		count++
	}

	return node
}

//尾部添加
func (c *CirCleList) Append(node *CirCleNode) bool {
	if node == nil {
		return false
	}
	if c.Count == 0 {
		c.Tail = node
		c.Head = node
		node.Next = nil
		node.Pre = node
	} else {
		node.Pre = c.Tail
		node.Next = nil
		c.Tail.Next = node
		c.Tail = node
	}
	c.Count++
	return true
}

func (c *CirCleList) Insert(index int, node *CirCleNode) bool {
	if node == nil || index > c.Count {
		return false
	}

	if index == c.Count {
		c.Append(node)
	} else {
		cirCleNode := c.GetIndex(index)

		node.Pre = cirCleNode.Pre
		node.Next = cirCleNode
		cirCleNode.Pre.Next = node
		cirCleNode.Pre = node
	}

	c.Count++

	return true

}

func (c *CirCleList) Delete(index int) bool {
	//index 从0开始
	if index > c.Count-1 {
		return false
	}
	//从头部删除
	cur := c.GetIndex(index)

	if index <= 0 {
		if c.Count == 1 {
			c.New()
		} else {
			c.Tail = cur.Next
		}
	} else {
		c.Tail = cur.Next
	}
	c.Count--
	return true
}

func (c *CirCleList) IsHead(node *CirCleNode) bool {
	return c.Head == node
}

func (c *CirCleList) IsTail(node *CirCleNode) bool {
	return c.Tail == node
}

//获取节点数据
func (c *CirCleNode) GetNodeData() interface{} {
	return c.Data
}

//获取下一个节点
func (c *CirCleNode) GetNext() *CirCleNode {
	return c.Next
}

//获取上一个节点
func (c *CirCleNode) GetPre() *CirCleNode {
	return c.Pre
}
