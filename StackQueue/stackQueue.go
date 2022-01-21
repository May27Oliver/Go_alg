package main

import "fmt"

/*
stack
last in first out 後進先出
*/
type Stack struct{
	items []int
}
//Push, add a value at the end
func (s *Stack) Push(val int){
	s.items = append(s.items,val)
}

//Pop will remove a value at the end
func (s *Stack) Pop()int{
	lastIndex := len(s.items)-1
	popIdx := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return popIdx
}

/*
queue
first in first out 先進先出
*/
type Queue struct{
	items []int
}
//Enqueue adds a value at the end
func (q *Queue) Enqueue(val int){
	q.items = append(q.items,val)
}

func (q *Queue) Dequeue()int{
	dequeueNum := q.items[0]
	q.items = q.items[1:]
	return dequeueNum
}

func main(){
	//Stack
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
}