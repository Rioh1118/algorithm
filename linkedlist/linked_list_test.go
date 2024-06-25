package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	t.Run("should empty", func(t *testing.T) {
		assert.Equal(t, true, list.IsEmpty())
	})
	t.Run("should size 0", func(t *testing.T) {
		assert.Equal(t, 0, list.Size())
	})
}

func TestPushFront(t *testing.T) {
	int_list := NewLinkedList[int]()
	for i := 0; i < 10; i++ {
		int_list.PushFront(i)
	}
	t.Run("should size 10", func(t *testing.T) {
		assert.Equal(t, 10, int_list.Size())
	})
	t.Run("should not empty", func(t *testing.T) {
		assert.Equal(t, false, int_list.IsEmpty())
	})
	t.Run("shold shape [9, 8, 7, 6, 5, 4, 3, 2, 1, 0 ]", func(t *testing.T) {
		assert.Equal(t, int_list.String(), "[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]")
	})

}
