package queue

import (
	"fmt"
	"log"
)

type Queue struct {
	arr      []int
	capacity int
	front    int
	rear     int
}

func NewQueue(arr []int, capacity int) *Queue {
	queue := &Queue{
		arr:      arr,
		capacity: capacity,
		front:    -1,
		rear:     -1,
	}

	return queue
}

func (q *Queue) isFull() bool {
	if q.front == 0 && q.rear == q.capacity-1 {
		return true
	}
	return false
}

func (q *Queue) isEmpty() bool {
	if q.front == -1 {
		return true
	}
	return false
}

func (q *Queue) EnQueue(n int) {
	if q.isFull() {
		log.Fatal("QUEUE IS FULL")
	} else {
		if q.front == -1 {
			q.front = 0
		}
		q.rear++
		q.arr = append(q.arr, n)
		fmt.Println("Inserted: ", n)
	}
}
func (q *Queue) DeQueue() {
	var e int
	if q.isEmpty() {
		log.Fatal("QUEUE IS EMPTY")
	} else {
		e = q.arr[q.front]
		if q.front >= q.rear {
			q.front = -1
			q.rear = -1
		} else {
			q.front++
		}
		q.arr = q.arr[1:]
		fmt.Println("Deleted: ", e)
	}
}
