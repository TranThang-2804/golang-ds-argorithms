package doublelinkedlist

import (
	"fmt"
	"slices"
)

type Comparator[T comparable] func(a, b T) int

// Node is a single element in a linked list.
type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

// DoubleLinkedList struct
type DoubleLinkedList[T comparable] struct {
	head *Node[T]
	last *Node[T]
	size int
}

// Create a new empty linked list
func New[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{head: nil, last: nil, size: 0}
}

// Append a new node to the end of linked list
func (l *DoubleLinkedList[T]) Append(items ...T) {
	for _, item := range items {
		newNode := &Node[T]{value: item, next: nil, prev: nil}
		if l.size == 0 {
			l.head = newNode
		} else {
			newNode.prev = l.last
			l.last.next = newNode
		}
		l.last = newNode
		l.size++
	}
}

// Append a new node to the beginning of linked list
func (l *DoubleLinkedList[T]) Prepend(items ...T) {
	for i := len(items) - 1; i >= 0; i-- {
		newNode := &Node[T]{value: items[i], next: nil, prev: nil}
		if l.size == 0 {
			l.last = newNode
		} else {
			l.head.prev = newNode
			newNode.next = l.head
		}
		l.head = newNode
		l.size++
	}
}

// Get the item of the link list at the specified
// err return true if the item is found else return false
func (l *DoubleLinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var t T
		return t, false
	}

	node := l.head
	for i := 0; i != index; i, node = i+1, node.next {
	}

	return node.value, true
}

// Remove the item of the link list
func (l *DoubleLinkedList[T]) Remove(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}

	// If the linked list has only 1 item
	if l.size == 1 {
		l.Clear()
		return true
	}

	// current node starts at index 1
	prevNode := l.head
	currentNode := l.head.next

	switch index {
	case 0:
		// If the index is 0, then we need to remove the head
		l.head = prevNode.next
		l.head.prev = nil
		prevNode = nil
		break
	case l.size - 1:
		// If the index is the last item
		l.last = l.last.prev
		l.last.next = nil
	default:
		// If the index is in the middle of the list
		for i := 1; i < index; i++ {
			prevNode = prevNode.next
			currentNode = currentNode.next
		}

		prevNode.next = currentNode.next
		prevNode.next.prev = prevNode
		currentNode = nil
		break
	}

	l.size--
	return true
}

// Check if the linked list contains the item
// return true if the item is found else return false
func (l *DoubleLinkedList[T]) Contains(item T) bool {
	for current := l.head; current != nil; current = current.next {
		if current.value == item {
			return true
		}
	}
	return false
}

// Return an array of all the items in the linked list
func (l *DoubleLinkedList[T]) GetAllNode() []T {
	var items []T
	for current := l.head; current != nil; current = current.next {
		items = append(items, current.value)
	}
	return items
}

// Get the size of the linked list
func (l *DoubleLinkedList[T]) GetSize() int {
	return l.size
}

// Check if the linked list is empty
func (l *DoubleLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Sort the linked list with the input is a compareFunction
func (l *DoubleLinkedList[T]) Sort(compareFunction Comparator[T]) {
	nodeList := l.GetAllNode()
	slices.SortFunc(nodeList, compareFunction)
	l.Clear()
	l.Append(nodeList...)
}

// Swap 2 items in the linked list
func (l *DoubleLinkedList[T]) Swap(i, j int) {
	if i < 0 || i >= l.size || j < 0 || j >= l.size {
		return
	}

	var nodeI, nodeJ *Node[T]

	for index, current := 0, l.head; current != nil; index, current = index+1, current.next {
		if index == i {
			nodeI = current
		}
		if index == j {
			nodeJ = current
		}
	}

	nodeI.value, nodeJ.value = nodeJ.value, nodeI.value
}

// Insert items at the specified index
// return true if the items are inserted successfully else return false
func (l *DoubleLinkedList[T]) Insert(index int, items ...T) bool {
	if index < 0 || index > l.size {
		return false
	}

	switch index {
	case 0:
		// If insrt at the beginning of the list
		l.Prepend(items...)
		break
	case l.size:
		// If insrt at the end of the list
		l.Append(items...)
		break
	default:
		// Insert in the middle of the list
		currentNode := l.head

		// Finding the node corresponding to the index
		for i := 1; i < index; i++ {
			currentNode = currentNode.next
		}

		// Inserting the items
		for _, item := range items {
    newNode := &Node[T]{value: item, next: currentNode.next, prev: currentNode}
			currentNode.next = newNode
			currentNode = newNode
			l.size++
		}
	}
	return true
}

// Update the value of the node at the specified index with the new value
// return true if the value is updated successfully else return false
func (l *DoubleLinkedList[T]) UpdateNodeValue(index int, item T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	current := l.head
	for i := 0; i != index; i, current = i+1, current.next {
	}

	current.value = item
	return true
}

// Go through the whole list and return the value as string
func (l *DoubleLinkedList[T]) String() string {
	str := "DoubleLinkedList\n"
	for currentItem := l.head; currentItem != nil; currentItem = currentItem.next {
		str += fmt.Sprintf("%v", currentItem.value)
	}
	return str
}

// Clear all item in the linked list
func (list *DoubleLinkedList[T]) Clear() {
	list.size = 0
	list.head = nil
	list.last = nil
}
