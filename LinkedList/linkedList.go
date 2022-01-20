package main

import "fmt"

/*
Linked List
Linked List串列和array一樣都是有序的存放元素的集合，
惟串列在記憶體中並不是連續放置，每個元素由一個存放元素本身的節點和一個指向下一個元素的pointer組成。

串列的好處在於添加或移除元素時不需要移動其他元素，但找尋串列中的某個元素必須從頭找起，陣列則可以直接找到元素。
*/
type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

//在串列中最前面插入元素
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

//刪除串列中的值
func (l *linkedList) deleteValue(value int) {
	//串列為空返回
	if l.length == 0{
		return
	}
	//數值與head.data相等
	if l.head.data == value{
		l.head = l.head.next
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
	l.length--
}

//在串列最後插入元素
func (l *linkedList) append(val int){
	pointer:=l.head
	for pointer.next != nil{
		pointer = pointer.next
	}
	pointer.next = &node{data:val}
	l.length++
}
//在任意位置插入元素
func (l *linkedList) insert(val, pos int){
	if pos < 2 {
		fmt.Println("欲插入位置小於0，直接加入串列頭部！")
		l.prepend(&node{data:val})
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
	pointer.next = &node{data:val,next:temp}
	l.length++
}

//返回元素在串列中的索引
func (l *linkedList) indexOf(el int)int{
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
func (l *linkedList) deleteAt(pos int){
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
	l.length--
}
//印出所有的串列元素
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d \n", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func main() {
	mylist := linkedList{}
	mylist.prepend(&node{data: 48}) //實體化一個node data為48然後把記憶體位置prepend進linkedList
	mylist.prepend(&node{data: 18})
	mylist.prepend(&node{data: 16})
	mylist.prepend(&node{data: 11})
	mylist.prepend(&node{data: 7})
	mylist.prepend(&node{data: 2})
	mylist.append(17)
	mylist.printListData()
	mylist.insert(5,43)
	mylist.printListData()
	fmt.Println("位置在：",mylist.indexOf(7))
	mylist.deleteAt(1)
	mylist.printListData()
}
