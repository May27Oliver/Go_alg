package main

import "fmt"

type Node struct{
	data int
	next *Node
	prev *Node
}

type LinkedList struct{
	head *Node
	tail *Node
	length int
}

//在串列中最前面插入元素
func (l *LinkedList) prepend(n *Node) {
	if l.head == nil{
		l.head = n
		l.tail = n
	}else{
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
	l.length++
	return
}

//刪除串列中的值
func (l *LinkedList) deleteValue(value int) {
	//串列為空返回
	if l.length == 0{
		return
	}
	//數值與head.data相等
	if l.head.data == value{
		l.head = l.head.next
		l.head.prev = nil
		l.length--
		return
	}
	previousToDelete := l.head
	//找尋node的next是否為我們要刪除的數值
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil{
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	previousToDelete.next.prev = previousToDelete
	l.length--
}

//在串列最後插入元素
func (l *LinkedList) append(val int){
	pointer:=l.tail
	pointer.next = &Node{data:val}
	pointer.next.prev = pointer
	l.tail = pointer.next
	l.length++
}
//在任意位置插入元素
func (l *LinkedList) insert(val, pos int){
	if pos < 2 {
		fmt.Println("欲插入位置小於0，直接加入串列頭部！")
		l.prepend(&Node{data:val})
		return
	} 
	if pos > l.length - 1 {
		fmt.Println("欲插入位置超過串列長度，直接加入串列最末位！")
		l.append(val)
		return
	}
	pointer := l.head
	listLen := 0
	for listLen < pos - 2 {
		pointer = pointer.next
		listLen++
	}
	fmt.Println("listLen",listLen,"pos",pos)
	temp := pointer.next
	pointer.next = &Node{data:val, next:temp, prev:pointer}
	temp.prev = pointer.next
	l.length++
}

//返回元素在串列中的索引
func (l *LinkedList) indexOf(el int)int{
	pointer := l.head
	var pos int
	for pos = 0 ;pos<l.length; pos++{
		if pointer.data == el{
			return pos
		}
		pointer = pointer.next
	}
	return -1
}

//刪除串列中特定位置元素
func (l *LinkedList) deleteAt(pos int){
	if pos > l.length-1{
		return
	}
	pointer:=l.head
	var	i int
	for i < pos-1{
		pointer= pointer.next
		i++
	}
	pointer.next = pointer.next.next
	pointer.next.prev = pointer
	l.length--
}
//印出所有的串列元素
func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d \n", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
} 


func main(){
	mylist := LinkedList{}
	mylist.prepend(&Node{data: 48}) //實體化一個node data為48然後把記憶體位置prepend進linkedList
	mylist.prepend(&Node{data: 18})
	mylist.prepend(&Node{data: 16})
	mylist.prepend(&Node{data: 11})
	mylist.prepend(&Node{data: 7})
	mylist.prepend(&Node{data: 2})
	mylist.append(17)
	mylist.printListData()
	mylist.insert(5,43)
	mylist.printListData()
	fmt.Println("位置在：",mylist.indexOf(7))
	mylist.deleteAt(1)
	mylist.printListData()
}