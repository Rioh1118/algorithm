package linkedlist

import (
	"testing"
)

// Helper function to compare slices
func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPushFront(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushFront(10)
	list.PushFront(20)
	list.PushFront(30)

	expected := []int{30, 20, 10}
	for i, v := range expected {
		value, err := list.ValueAt(i)
		if err != Nil || value != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, value)
		}
	}
}

func TestPushBack(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	expected := []int{10, 20, 30}
	for i, v := range expected {
		value, err := list.ValueAt(i)
		if err != Nil || value != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, value)
		}
	}
}

func TestPopFront(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushFront(10)
	list.PushFront(20)
	list.PushFront(30)

	value, err := list.PopFront()
	if err != Nil || value != 30 {
		t.Errorf("expected value 30, got %d", value)
	}

	expected := []int{20, 10}
	for i, v := range expected {
		value, err := list.ValueAt(i)
		if err != Nil || value != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, value)
		}
	}
}

func TestPopBack(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	value, err := list.PopBack()
	if err != Nil || value != 30 {
		t.Errorf("expected value 30, got %d", value)
	}

	expected := []int{10, 20}
	for i, v := range expected {
		value, err := list.ValueAt(i)
		if err != Nil || value != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, value)
		}
	}
}

func TestValueAt(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	value, err := list.ValueAt(1)
	if err != Nil || value != 20 {
		t.Errorf("expected value 20 at index 1, got %d", value)
	}

	_, err = list.ValueAt(3)
	if err != IndexOutOfRange {
		t.Errorf("expected IndexOutOfRange error, got %v", err)
	}
}

func TestFront(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	value, err := list.Front()
	if err != Nil || value != 10 {
		t.Errorf("expected front value 10, got %d", value)
	}
}

func TestBack(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	value, err := list.Back()
	if err != Nil || value != 30 {
		t.Errorf("expected back value 30, got %d", value)
	}
}

func TestInsert(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	err := list.Insert(1, 25)
	if err != Nil {
		t.Errorf("expected no error, got %v", err)
	}

	expected := []int{10, 25, 20, 30}
	for i, v := range expected {
		value, err := list.ValueAt(i)
		if err != Nil || value != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, value)
		}
	}
}

func TestErase(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	err := list.erase(1)
	if err != Nil {
		t.Errorf("expected no error, got %v", err)
	}

	expected := []int{10, 30}
	for i, v := range expected {
		value, err := list.ValueAt(i)
		if err != Nil || value != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, value)
		}
	}
}

func TestValueNFromEnd(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)
	list.PushBack(40)
	list.PushBack(50)

	values, err := list.ValueNFromEnd(2)
	expected := []int{30, 40, 50}
	if err != Nil || !equal(values, expected) {
		t.Errorf("expected values %v, got %v", expected, values)
	}
}

func TestReverse(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)
	list.PushBack(40)
	list.PushBack(50)

	reversed := list.Reverse()

	expected := []int{50, 40, 30, 20, 10}
	for i, v := range expected {
		value, err := reversed.ValueAt(i)
		if err != Nil || value != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, value)
		}
	}
}

type stringerValue string

func (s stringerValue) String() string {
	return string(s)
}
