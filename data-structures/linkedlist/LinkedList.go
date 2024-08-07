package linkedlist

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

func (this *Node) Next() *Node {
	return this.next
}

func (this *Node) Prev() *Node {
	return this.prev
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{length: 0}
}

func (this *LinkedList) GetLength() int {
	return this.length
}

func (this *LinkedList) GetHead() *Node {
	return this.head
}

func (this *LinkedList) GetLast() *Node {
	return this.Get(this.GetLength() - 1)
}

func (this *LinkedList) Get(index int) *Node {
	var currentIndex = 0
	var current = this.head
	for currentIndex < index {
		current = current.next
		currentIndex++
	}

	return current
}

func (this *LinkedList) Push(value int) {
	var isListEmpty = this.head == nil
	if isListEmpty {
		this.head = &Node{value, nil, nil}
	} else {
		var oldLastNode = this.GetLast()
		oldLastNode.next = &Node{value, nil, oldLastNode}
	}

	this.length++
}

func (this *LinkedList) Unshift(value int) {
	var isListEmpty = this.head == nil
	if isListEmpty {
		this.head = &Node{value, nil, nil}
	} else {
		var oldHead = this.head
		this.head = &Node{value, oldHead, nil}
		oldHead.prev = this.head
	}
	this.length++
}

func (this *LinkedList) Insert(value int, afterIndex int) {
	var current = this.Get(afterIndex)
	var isCurrentNodeLast = current.next == nil

	if isCurrentNodeLast {
		this.Push(value)
	} else {
		var oldNext = current.next
		current.next = &Node{value, oldNext, current}
		oldNext.prev = current.next
		this.length++
	}
}

func (this *LinkedList) Pop() {
	var last = this.GetLast()
	var prev = last.prev

	prev.next = nil
	last.prev = nil

	this.length--
}

func (this *LinkedList) Shift() {
	var oldHead = this.head

	this.head = oldHead.next
	this.head.prev = nil

	oldHead.next = nil

	this.length--
}

func (this *LinkedList) Delete(index int) {
	var nodeToDelete = this.Get(index)
	var prev = nodeToDelete.prev
	var next = nodeToDelete.next

	if prev == nil {
		this.Shift()
	} else if next == nil {
		this.Pop()
	} else {
		nodeToDelete.prev = nil
		nodeToDelete.next = nil
		prev.next = next
		next.prev = prev
		this.length--
	}
}

func (this *LinkedList) ToArray() []int {
	var result = make([]int, 0, this.GetLength())
	current := this.head
	for current != nil {
		result = append(result, current.Value)
		current = current.next
	}
	return result
}

func (this *LinkedList) ToReverseArray() []int {
	var result = make([]int, 0, this.GetLength())
	current := this.GetLast()
	for current != nil {
		result = append(result, current.Value)
		current = current.prev
	}
	return result
}
