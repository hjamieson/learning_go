package main

import (
	"fmt"
	"io"
)

type List[T comparable] struct {
	head *Element[T]
	tail *Element[T]
}

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

func (l *List[T]) Add(value T) {
	e := &Element[T]{value: value, next: nil}
	if l.head == nil {
		l.head = e
		l.tail = e
		return
	}
	l.tail.next = e
	l.tail = e
}

func (l *List[T]) Insert(value T, index int) {
	// replacing head?
	if l.head == nil {
		l.Add(value)
		return
	}
	e := &Element[T]{value, nil}
	if index == 0 {
		e.next = l.head
		l.head = e
	} else {
		p := l.head
		ix := 0
		for {
			if ix < index-1 {
				ix++
				p = p.next
			} else {
				break
			}
		}
		e.next = p.next
		p.next = e
	}
}

func (l *List[T]) Index(value T) int {
	ix := 0
	for curNode := l.head; curNode != nil; curNode = curNode.next {
		if curNode.value == value {
			return ix
		}
		ix++
	}
	return -1
}

func NewList[T comparable]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Dump(w io.Writer) {
	p := l.head
	for {
		if p != nil {
			io.WriteString(w, fmt.Sprintf("%v\n", p.value))
			p = p.next
		} else {
			break
		}
	}
}
