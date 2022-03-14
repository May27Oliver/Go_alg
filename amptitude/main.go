package main

import (
	"fmt"
)

func main(){
	a:=[]int{5,3,6,1,3}
	s:=Solution(a,2)
	fmt.Println(s)
}

type MaxHeap struct{
	arr []int
}

func (h *MaxHeap) Insert(value int){
	h.arr = append(h.arr,value)
	h.maxHeapifyUp(len(h.arr)-1)
}

func (h *MaxHeap) maxHeapifyUp(i int) {
	for h.arr[i] > h.arr[parent(i)] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func Solution(A []int, K int) int {
	// write your code in Go 1.4
	/*
	思路：
	用MaxHeap排序刪除連續K值後的陣列，紀錄最小的amptitude，存放入一個陣列，最後找出陣列最小的amptitude
	*/
	var counter []int//紀錄各個amptitude
	var record [][]int//紀錄各個arr
	var start,end = 0, K
	for end <= len(A){
		//Heap裡的arr紀錄的是各個arr去掉連續K個值後，最大和次大的數的差距
		m:=&MaxHeap{}
		for i := 0;i < len(A);i++ {
			if i < start || i > end - 1 {
				m.Insert(A[i])
			}
		}
		//Heap node與左右子數相減，紀錄最小的值，最小的存進counter內
		left_amptitude:=m.arr[0]-m.arr[1]
		right_amptitude:=m.arr[0]-m.arr[2]
		if left_amptitude < right_amptitude{
			counter = append(counter, left_amptitude)
		}else{
			counter = append(counter,right_amptitude)
		}
		record = append(record,m.arr)
		end++
		start++
	}
	fmt.Println(counter)
	minAmptitude := 1000000//先假設minAmptitude為一無窮大數
	KstartAt := 0
	for i,v:=range counter{
		if minAmptitude > v{
			minAmptitude = v
			KstartAt = i
		}
	}
	fmt.Println(record[KstartAt])
	//找出counter內最小的值，紀錄其key值，就是K值
	return minAmptitude
}