package main

import (
	"fmt"
	"io"
)

type TreeCompare[T comparable] func(T, T) int

type Tree[T comparable] struct {
	root *Node[T]
	f    TreeCompare[T]
}

type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func (n *Node[T]) addNode(val T, f TreeCompare[T]) {
	switch o := f(n.value, val); {
	case o < 0:
		if n.left == nil {
			n.left = &Node[T]{value: val}
		} else {
			n.left.addNode(val, f)
		}
	case o > 0:
		if n.right == nil {
			n.right = &Node[T]{value: val}
		} else {
			n.right.addNode(val, f)
		}
	}
}

func NewTree[T comparable](f TreeCompare[T]) *Tree[T] {
	return &Tree[T]{
		nil,
		f,
	}
}

func (t *Tree[T]) Add(val T) {
	switch t.root {
	case nil:
		t.root = &Node[T]{value: val}
	default:
		t.root.addNode(val, t.f)
	}
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(v)
}

func (n *Node[T]) Contains(v T) bool {
	if n == nil {
		return false
	}
	if n.value == v {
		return true
	}
	return n.left.Contains(v) || n.right.Contains(v)
}

func (n *Node[T]) Dump(w io.Writer) {
	if n.right != nil {
		n.right.Dump(w)
	}
	io.WriteString(w, fmt.Sprintf("%v\n", n.value))
	if n.left != nil {
		n.left.Dump(w)
	}
}

func (t *Tree[T]) Dump(w io.Writer) {
	// process right
	t.root.right.Dump(w)
	// process this
	io.WriteString(w, fmt.Sprintf("%v\n", t.root.value))
	// process left
	t.root.left.Dump(w)
}
