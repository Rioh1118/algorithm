package linkedlist

import (
	"fmt"
	"strings"
)

type Error byte

const (
	IndexOutOfRange Error = iota
	NotExistValue
	UnKnown
	Nil
)

func (e Error) Error() string {
	switch e {
	case IndexOutOfRange:
		return "Index out of range"
	case NotExistValue:
		return "Value does not exist"
	case UnKnown:
		return "Unknown error"
	default:
		return "Unknown error"
	}
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (list *LinkedList[T]) Size() int {
	return list.length
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *LinkedList[T]) ValueAt(index int) (T, Error) {
	if index < 0 || index >= list.length {
		var v T
		return v, IndexOutOfRange
	} else {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current.value, Nil
	}

}

func (list *LinkedList[T]) PushFront(value T) {
	if list.IsEmpty() {
		list.head = &Node[T]{value: value, next: nil}
		list.tail = list.head
		list.length++
	} else {
		node := &Node[T]{value: value, next: list.head}
		list.head = node
		list.length++
	}
}

func (list *LinkedList[T]) PopFront() (T, Error) {
	if list.IsEmpty() {
		var v T
		return v, NotExistValue
	}
	value := list.head.value
	list.head = list.head.next
	list.length--
	return value, Nil
}

func (list *LinkedList[T]) PushBack(value T) {
	if list.IsEmpty() {
		list.PushFront(value)
		list.length++
	} else {
		node := &Node[T]{value: value, next: nil}
		list.tail.next = node
		list.tail = node
		list.length++
	}
}

func (list *LinkedList[T]) PopBack() (T, Error) {
	if list.IsEmpty() {
		var v T
		return v, NotExistValue
	} else if list.Size() == 1 {
		return list.PopFront()
	}
	for current := list.head; current.next != nil; current = current.next {
		if current.next == list.tail {
			value := list.tail.value
			list.tail = current
			current.next = nil
			list.length--
			return value, Nil
		}
	}
	var v T
	return v, UnKnown
}

func (list *LinkedList[T]) Front() (T, Error) {
	if list.IsEmpty() {
		var v T
		return v, NotExistValue
	}
	return list.head.value, Nil
}
func (list *LinkedList[T]) Back() (T, Error) {
	if list.IsEmpty() {
		var v T
		return v, NotExistValue
	}
	return list.tail.value, Nil
}

func (list *LinkedList[T]) Insert(index int, value T) Error {
	if index < 0 || index >= list.length {
		return IndexOutOfRange
	}
	current := list.head
	next := current.next
	if index == 0 {
		list.PushFront(value)
	}
	for i := 0; i < index; i++ {
		if i == index-1 {
			node := &Node[T]{value: value, next: next}
			current.next = node
			list.length++
			break
		}
		current = next
		next = current.next
	}
	return Nil
}

func (list *LinkedList[T]) erase(index int) Error {
	if 0 > index || index >= list.length {
		return IndexOutOfRange
	}
	if index == 0 {
		_, err := list.PopFront()
		return err
	}
	current := list.head
	next := current.next
	for i := 0; i < index; i++ {
		if i == index-1 {
			current.next = next.next
			list.length--
			return Nil
		}
	}
	return UnKnown
}

func (list *LinkedList[T]) ValueNFromEnd(n int) ([]T, Error) {
	if n < 0 || n >= list.length {
		return nil, IndexOutOfRange
	}
	result := make([]T, 100)
	current := list.getNode(n)
	for current.next != nil {
		result = append(result, current.value)
		current = current.next
	}
	result = append(result, current.value)

	return result, Nil
}

func (list *LinkedList[T]) getNode(index int) *Node[T] {
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}

func (list *LinkedList[T]) Reverse() LinkedList[T] {
	if list.IsEmpty() {
		return *list
	}
	memo := make(map[int]T, list.length)
	current := list.head
	for i := 0; i < list.length; i++ {
		memo[i] = current.value
		current = current.next
	}
	reversed := NewLinkedList[T]()
	for i := list.length - 1; i >= 0; i-- {
		reversed.PushBack(memo[i])
	}
	return *reversed
}

func (list *LinkedList[T]) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "[")
	if list.Size() == 0 {
		return "[]"
	}
	current := list.head
	for current.next != nil {
		fmt.Fprintf(&sb, "%v, ", current.value)
		current = current.next
	}
	fmt.Fprintf(&sb, "%v", current.value)

	fmt.Fprintf(&sb, "]")

	return sb.String()
}
