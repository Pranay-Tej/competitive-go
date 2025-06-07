package linkedlist

import "fmt"

func (ll *LinkedList[T]) Push(value T) {
	newNode := Node[T]{Value: value, Next: nil}
	if ll.Head == nil {
		ll.Head = &newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &newNode
}

func (ll *LinkedList[T]) ToSlice() []T {
	var result = []T{}
	current := ll.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func NewLinkedList[T any](slice []T) *LinkedList[T] {
	ll := LinkedList[T]{}
	for _, value := range slice {
		ll.Push(value)
	}
	return &ll
}

func (ll *LinkedList[T]) GetTail() *Node[T] {
	if ll.Head == nil {
		return nil
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

func (ll *LinkedList[T]) Pop() (T, bool) {
	if ll.Head == nil {
		var zero T
		return zero, false
	}
	if ll.Head.Next == nil {
		val := ll.Head.Value
		ll.Head = nil
		return val, true
	}
	var prev *Node[T] = nil
	current := ll.Head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	prev.Next = nil
	return current.Value, true
}

func (ll *LinkedList[T]) GetAtIndex(index int) (*Node[T], error) {
	if index < 0 {
		return nil, fmt.Errorf("invalid index")
	}
	current := ll.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return nil, fmt.Errorf("index out of range")
		}
		current = current.Next
	}
	return current, nil
}

func (ll *LinkedList[T]) Reverse() {
	if ll.Head == nil {
		return
	}

	var prev *Node[T] = nil
	current := ll.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev
}
