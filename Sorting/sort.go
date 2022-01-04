package main

import "fmt"

var x = []int{4,7,1,2,5,3,8,10,33,6,77,100}

func main() {
	y:=bubbleSort(x)
	fmt.Print(y)
}
/*
	BubbleSort
	由後到前兩兩比較，較小的數字往前排，由此排序n平方次。
	1.設定一個pointer，len(arr)，loop讓arr[len(arr)]往前跑去比較與arr[len(arr)-1]大小
	2.第一次迴圈找出陣列第一個值，也就是最小值，第二次迴圈找出陣列第二個值次小值。
	3.所以需要兩個迴圈一個由最末位往前前後比較，一個遍歷所有迴圈讓每個值都排序。
	Bubble Sort的BigO是n2
	因為每一位都要由1到n比對一次，所以(n-1)+(n-2)+(n-3)+(n-4)+.....+0+1=((n-1)*n)/2 => n2( 常數省略、小數省略 )
	最糟的狀況BubbleSort要O(n^2)
	最好的狀況(陣列本身就排好)，則只需要O(n)
*/
func bubbleSort(arr []int)[]int{
	var temp int
	for _,_=range arr{
		for i:=len(arr)-2; i >= 0; i-- {
			if arr[i+1] < arr[i]{
				temp = arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
			}
		}
	}
	return arr
}

/*
Insertion Sort插入排序法
*/ 
func insertionSort(arr []int)[]int{
	for i:=1;i<len(arr);i++ {
		temp := arr[i]
		j := i-1
		for j >= 0 && temp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}