package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	q []int
}

func (q *Queue) Push(val int) {
	q.q = append(q.q, val)
}
func (q *Queue) Pop() (int, error) {
	if len(q.q) > 0 {
		val := q.q[0]
		q.q = q.q[1:]
		return val, nil
	}

	return -1, errors.New("queue is empty")
}
func main() {
	queue := &Queue{}

	queue.Push(10)
	queue.Push(20)
	queue.Push(30)
	queue.Push(40)
	queue.Push(50)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

}
