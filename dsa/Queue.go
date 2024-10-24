package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	q []int
}

func (q *Queue) push(val int) {
	q.q = append(q.q, val)
}
func (q *Queue) pop() (int, error) {
	if len(q.q) > 0 {
		val := q.q[0]
		q.q = q.q[1:]
		return val, nil
	}

	return -1, errors.New("queue is empty")
}
func main() {
	queue := &Queue{}

	queue.push(10)
	queue.push(20)
	queue.push(30)
	queue.push(40)
	queue.push(50)
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())

}
