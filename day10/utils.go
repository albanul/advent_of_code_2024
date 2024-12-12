package queue

import "errors"

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.elements) == 0 {
		var zeroValue T
		return zeroValue, errors.New("empty queue")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}
