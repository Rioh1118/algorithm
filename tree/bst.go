package tree

import (
	"fmt"
	"strings"
)

type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}
type Comparable interface {
	comparable
	Ordered
}

type Node[T Comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T Comparable](value T) *Node[T] {
	return &Node[T]{value: value, left: nil, right: nil}
}

func (root *Node[T]) String() string {
	var sb strings.Builder
	root.buildString(&sb, "", true)
	return sb.String()
}

// buildStringはツリー構造を構築するヘルパーメソッドです。
func (root *Node[T]) buildString(sb *strings.Builder, prefix string, isTail bool) {
	if root.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		root.right.buildString(sb, newPrefix, false)
	}

	sb.WriteString(prefix)
	if isTail {
		sb.WriteString("└── ")
	} else {
		sb.WriteString("┌── ")
	}
	sb.WriteString(fmt.Sprintf("%v\n", root.value))

	if root.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		root.left.buildString(sb, newPrefix, true)
	}
}

func (root *Node[T]) Insert(value T) {
	if root == nil {
		return
	}

	if value < root.value {
		if root.left == nil {
			root.left = NewNode(value)
		} else {
			root.left.Insert(value)
		}
	} else if value > root.value {
		if root.right == nil {
			root.right = NewNode(value)
		} else {
			root.right.Insert(value)
		}
	}

	// if root == nil {
	// 	root = NewNode(value)
	// } else if value < root.value {
	// 	root.left.Insert(value)
	// } else if value > root.value {
	// 	root.right.Insert(value)
	// } else {
	// 	return
	// }
}

func (root *Node[T]) Get_node_count() int {
	if root == nil {
		return 0
	}

	left_count := root.left.Get_node_count()
	right_count := root.right.Get_node_count()

	return 1 + left_count + right_count
}

func (root *Node[T]) Print_values() {
	if root == nil {
		return
	}

	root.left.Print_values()
	fmt.Println(root.value)
	root.right.Print_values()
}

func (root *Node[T]) Is_in_tree(value T) bool {
	if root == nil {
		return false
	}
	if value == root.value {
		return true
	}
	return root.left.Is_in_tree(value) ||
	root.right.Is_in_tree(value)
}

func (root *Node[T]) get_min() T {
	if root.left == nil {
		return root.value
	}
	return root.left.get_min()
}

func (root *Node[T]) get_max() T {
	if root.right == nil {
		return root.value
	}
	return root.right.get_max()
}

func (root *Node[T]) delete(value T) {
	if root == nil {
		return
	}
	if value < root.value {
		root.left.delete(value)
	}
	if value > root.value {
		root.right.delete(value)
	}
	if value == root.value {
		if root.left == nil && root.right == nil {
			root = nil
			return
		} else if root.left == nil && root.right != nil {
			root.value = root.right.value
			root.right = nil
		} else if root.left != nil && root.right == nil {
			root.value = root.left.value
			root.left = nil
		} else {
			right_min := root.right.get_min()
			root.value = right_min
			root.right.delete(right_min)
		}
	}
}
