package queue

import (
		"github.com/olamiko/stack"
		)

type Queue struct{
	StackOne stack.Stack
	StackTwo stack.Stack
}

func (q *Queue) Enqueue(item interface{}){
	q.StackOne.Push(item)
}

func (q *Queue) Dequeue() interface{}{

	if q.StackTwo.IsEmpty(){
		for !q.StackOne.IsEmpty() {
			q.StackTwo.Push(q.StackOne.Pop())
		}
	}

	return q.StackTwo.Pop()
}

func (q Queue) IsEmpty() bool{
	return q.StackOne.IsEmpty() && q.StackTwo.IsEmpty()
}

func (q Queue) Size() int{
	return q.StackOne.Size() + q.StackTwo.Size()
}
