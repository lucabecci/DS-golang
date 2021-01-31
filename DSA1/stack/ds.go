package stack

import (
	"fmt"
	"log"
)

type Stack struct {
	arr      []int
	capacity int
	top      int
}

func NewStack(array []int, capacity int) *Stack {
	stack := &Stack{
		arr:      array,
		capacity: capacity,
		top:      -1,
	}

	return stack
}

func (s *Stack) Push(n int) {
	if s.isFull() {
		log.Panic("The Stack Is Full")
	}
	fmt.Println("Inserting: ", n)
	s.arr = append(s.arr, n)
	s.top++
}

func (s *Stack) Pop() {
	if s.isEmpty() {
		log.Panic("The Stack Is Empty")
	}
	s.arr = s.arr[:len(s.arr)-1]
	s.top--
}

func (s *Stack) size() int {
	return s.top + 1
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) isFull() bool {
	return s.top == s.capacity-1
}

func (s *Stack) PrintStack() {
	for _, num := range s.arr {
		fmt.Println("Num:", num)
	}
}
