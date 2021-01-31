package queue

import "fmt"

func Run() {
	var arr []int
	queue := Queue{
		arr:      arr,
		capacity: 5,
	}

	queue.EnQueue(4)
	queue.EnQueue(7)
	queue.EnQueue(3)
	queue.DeQueue()
	queue.EnQueue(9)
	queue.DeQueue()
	queue.EnQueue(22)
	queue.DeQueue()
	fmt.Println(queue.arr)
}
