package queue

import "fmt"

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) PrintQueue() {
	fmt.Println("Queue:")
	for _, item := range q.items {
		fmt.Println(item)
	}
}
