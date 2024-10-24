package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	st []int
}

func (s *Stack) push(val int) {
	s.st = append(s.st, val)
}

func (s *Stack) pop() (int, error) {
	if len(s.st) > 0 {
		val := s.st[len(s.st)-1]
		s.st = s.st[0 : len(s.st)-1]
		return val, nil
	}
	return -1, errors.New("stack is Empty")
}
func (s *Stack) IsEmpty() bool {
	if len(s.st) == 0 {
		return true
	}
	return false
}
func main() {
	st := &Stack{}
	st.push(10)
	st.push(20)
	st.push(30)
	fmt.Println(st.pop())
	fmt.Println(st.IsEmpty())
}
