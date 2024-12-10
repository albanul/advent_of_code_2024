package utils

import "errors"

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("empty queue")
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func NewQueue() *Queue {
	return &Queue{}
}
