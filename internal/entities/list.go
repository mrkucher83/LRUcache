package entities

type List struct {
	FirstNode *Node
	LastNode  *Node
	Amount    int
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
	return l.Amount
}

func (l *List) PushFront(v interface{}) {
	node := &Node{v, l.FirstNode, nil}
	if l.Len() == 0 {
		l.LastNode = node
	} else {
		l.FirstNode.Prev = node
	}

	l.FirstNode = node
	l.Amount++
}

func (l *List) PushBack(v interface{}) {
	node := &Node{v, nil, l.LastNode}
	if l.Len() == 0 {
		l.FirstNode = node
	} else {
		l.LastNode.Next = node
	}

	l.LastNode = node
	l.Amount++
}

func (l *List) Remove(node *Node) {
	if node.Prev == nil {
		l.FirstNode = node.Next
	} else {
		node.Prev.Next = node.Next
	}

	if node.Next == nil {
		l.LastNode = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}

	l.Amount--
}

func (l *List) MoveFront(node *Node) {
	if l.FirstNode == node {
		return
	}

	if l.LastNode == node {
		l.LastNode = node.Prev
		node.Prev.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	l.FirstNode.Prev = node
	node.Prev = nil
	node.Next = l.FirstNode
	l.FirstNode = node
}
