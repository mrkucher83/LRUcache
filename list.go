package LRUcache

type List struct {
	firstNode *Node
	lastNode  *Node
	amount    int
}

type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

func NewList() *List {
	return new(List)
}

func (l *List) Len() int {
	return l.amount
}

func (l *List) PushFront(v interface{}) *Node {
	node := &Node{v, l.firstNode, nil}
	if l.Len() == 0 {
		l.lastNode = node
	} else {
		l.firstNode.Prev = node
	}

	l.firstNode = node
	l.amount++

	return node
}

func (l *List) PushBack(v interface{}) {
	node := &Node{v, nil, l.lastNode}
	if l.Len() == 0 {
		l.firstNode = node
	} else {
		l.lastNode.Next = node
	}

	l.lastNode = node
	l.amount++
}

func (l *List) Remove(node *Node) {
	if node.Prev == nil {
		l.firstNode = node.Next
	} else {
		node.Prev.Next = node.Next
	}

	if node.Next == nil {
		l.lastNode = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}

	l.amount--
}

func (l *List) MoveToFront(node *Node) {
	if l.firstNode == node {
		return
	}

	if l.lastNode == node {
		l.lastNode = node.Prev
		node.Prev.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	l.firstNode.Prev = node
	node.Prev = nil
	node.Next = l.firstNode
	l.firstNode = node
}

func (l *List) First() *Node {
	return l.firstNode
}

func (l *List) Last() *Node {
	return l.lastNode
}
