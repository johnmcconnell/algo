package lnklst

type List interface {
	Get(int) interface{}
	Append(interface{})
	Size() int
}

type SingleLinkedList struct {
	head *Node
}

type Node struct {
	next *Node
	element interface{}
}

func NewSingleLinkedList() *SingleLinkedList {
	lst := SingleLinkedList{}
	return &lst
}

func newNode(e interface{}) *Node {
	n := Node{
		next: nil,
		element: e,
	}
	return &n
}

func (lst *SingleLinkedList) Append(e interface{}) {
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
