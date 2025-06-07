package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedListPush(t *testing.T) {
	ll := LinkedList[int]{}
	llArr := ll.ToSlice()
	if len(llArr) != 0 {
		t.Errorf("got %v", llArr)
	}
	ll.Push(5)
	llArr = ll.ToSlice()
	if !reflect.DeepEqual(llArr, []int{5}) {
		t.Errorf("got %v", llArr)
	}
	ll.Push(6)
	llArr = ll.ToSlice()
	if !reflect.DeepEqual(llArr, []int{5, 6}) {
		t.Errorf("got %v", llArr)
	}
}

func TestLinkedListPop(t *testing.T) {
	ll2 := NewLinkedList([]string{"a", "b", "c"})
	ll2Arr := ll2.ToSlice()
	if !reflect.DeepEqual(ll2Arr, []string{"a", "b", "c"}) {
		t.Errorf("got %v", ll2Arr)
	}
	poppedValue, ok := ll2.Pop()
	ll2Arr = ll2.ToSlice()
	if !ok || poppedValue != "c" || !reflect.DeepEqual(ll2Arr, []string{"a", "b"}) {
		t.Errorf("got %v", ll2Arr)
	}
	poppedValue, ok = ll2.Pop()
	ll2Arr = ll2.ToSlice()
	if !ok || poppedValue != "b" || !reflect.DeepEqual(ll2Arr, []string{"a"}) {
		t.Errorf("got %v", ll2Arr)
	}
	poppedValue, ok = ll2.Pop()
	ll2Arr = ll2.ToSlice()
	if !ok || poppedValue != "a" || !reflect.DeepEqual(ll2Arr, []string{}) {
		t.Errorf("got %v", ll2Arr)
	}
	poppedValue, ok = ll2.Pop()
	if ok {
		t.Errorf("got %v", ll2Arr)
	}
}

func TestLinkedListReverse(t *testing.T) {
	ll := NewLinkedList([]int{1, 2, 3, 4, 5})
	ll.Reverse()
	llArr := ll.ToSlice()
	if !reflect.DeepEqual(llArr, []int{5, 4, 3, 2, 1}) {
		t.Errorf("got %v", llArr)
	}
}
