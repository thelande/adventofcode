package queue

import "fmt"

var ErrQueueEmpty = fmt.Errorf("there are no elements in the queue")

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	head, tail *node[T]
	length     uint64
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Len() uint64 {
	return q.length
}

func (q *Queue[T]) Push(value T) error {
	node := &node[T]{
		value: value,
		next:  nil,
	}

	if q.tail != nil {
		q.tail.next = node
	} else if q.head == nil {
		q.head = node
	}

	q.tail = node
	q.length++

	return nil
}

func (q *Queue[T]) Pop() (T, error) {
	if q.head == nil {
		var empty T
		return empty, ErrQueueEmpty
	}

	if q.head == q.tail {
		q.tail = nil
	}

	oldHead := q.head
	q.head = oldHead.next
	q.length--
	return oldHead.value, nil
}

func (q *Queue[T]) Empty() bool {
	return q.length == 0
}
