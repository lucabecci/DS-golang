package stack

func Run() {
	var arr []int
	stack := NewStack(arr, 8)

	stack.Push(33)
	stack.Push(22)
	stack.Push(11)
	stack.Push(55)
	stack.Pop()
	stack.Push(88)

	stack.PrintStack()
}
