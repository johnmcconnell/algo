package lnklst

type List interface {
	Get(int) interface{}
	Add(interface{})
	Del(int) interface{}
	Size() int
	Rev()
}

type SingleLinkedList struct {
	head *Node
}

type Node struct {
	next    *Node
	element interface{}
}

func NewSingleLinkedList() *SingleLinkedList {
	lst := SingleLinkedList{}
	return &lst
}

func newNode(e interface{}) *Node {
	n := Node{
		next:    nil,
		element: e,
	}
	return &n
}

func revRec(n *Node) {
	if n.next == nil {
	}

	revRec(n.next)
}

func (lst *SingleLinkedList) Del(I int) interface{} {
	if I == 0 {
		r := lst.head.next
		lst.head = r
	}

	n := lst.head

	i := 1
	for i < I {
		n = n.next
	}

	r := n.next.next
	d := n.next.element
	n.next = r

	return d
}

func (lst *SingleLinkedList) Add(e interface{}) {
	if lst.head == nil {
		lst.head = newNode(e)
		return
	}

	tail := lst.head

	for tail.next != nil {
		tail = tail.next
	}

	tail.next = newNode(e)
}

func (lst *SingleLinkedList) Get(e int) interface{} {
	index := 0
	curr := lst.head

	for index < e {
		curr = curr.next
		index += 1
	}

	return curr.element
}
