package main

import "fmt"

var x = []int{4, 7, 1, 2, 5, 3, 8, 10, 33, 6, 77, 100, 101}

func main() {
	y := mergeSort(x)
	fmt.Print(y)
}

/* n^2 系列*/
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
func bubbleSort(arr []int) []int {
	var j = 0
	for j < len(arr) {
		for i := len(arr) - 2; i >= 0; i-- {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		j++
	}
	return arr
}

/*
	Insertion Sort插入排序法
	由前到後，排列第i項往前到0位的順序
	for i loop to len(arr)-1
	j = i - 1
	then sort j to 0 position

	Big O notation n*(n-1)/2
	複雜度n^2
*/
func insertionSort(arr []int) []int {
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			j--
		}
	}
	return arr
}

/*
	Selection Sort
	The principle of selection sort is - select the smallest value in unsorted array,
	and swap it with the left most value in this unsorted array.
	在陣列中找到最小值，依序排入index 0 ~ len(arr)-1

	selection sort一樣worst case 是 O(n^2), best case 也是 O(n^2)
*/
func selectionSort(arr []int) []int {
	var n = len(arr)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

/*
	Merge sort
	merge sort是很經典的"divide and merge"案例。
	merge sort利用拆分與合併兩個排列好的陣列兩個優點實現O(n logn n)的時間複雜度
	原理就是把未排序的大陣列切成以一個元素為單位的小陣列，再將之一個一個合併回排列好的大陣列。
	其實你會發現，merge sort就是Preorder Traversal。
	如果今天陣列的長度為n，要拆分log(2,n)次，且每一層要排序n次，所以mergesort的複雜度是n*log(2,n)
*/

func mergeSort(arr []int) []int {
		fmt.Println("拆分 arr",arr)
	//拆分
	if len(arr) > 1 {
		midIdx:=len(arr) / 2
		arr_left := mergeSort(arr[:midIdx])
		arr_right := mergeSort(arr[midIdx:])
		//合併
		return merge(arr_left,arr_right)
	}
	return arr
}

func merge(left_arr,right_arr []int)[]int{
	fmt.Println("合併left_arr",left_arr,"right_arr",right_arr)
	final:= []int{}
	i:=0
	j:=0
	for i < len(left_arr) && j < len(right_arr){
		if left_arr[i] < right_arr[j]{
			final = append(final, left_arr[i])
			i++
		} else {
			final = append(final,right_arr[j])
			j++
		}
	}
	fmt.Println("合併run after for",final)
	for ; i < len(left_arr); i++ {
			final = append(final, left_arr[i])
	}
	for ; j < len(right_arr); j++ {
			final = append(final, right_arr[j])
	}
	fmt.Println("合併run after append",final)
	return final      
}
/* log
拆分 arr [4 7 1 2 5 3 8 10 33 6 77 100 101]
拆分 arr [4 7 1 2 5 3]
拆分 arr [4 7 1]
拆分 arr [4]
拆分 arr [7 1]
拆分 arr [7]
拆分 arr [1]
合併left_arr [7] right_arr [1]
合併run after for [1]
合併run after append [1 7]
合併left_arr [4] right_arr [1 7]
合併run after for [1 4]
合併run after append [1 4 7]
拆分 arr [2 5 3]
拆分 arr [2]
拆分 arr [5 3]
拆分 arr [5]
拆分 arr [3]
合併left_arr [5] right_arr [3]
合併run after for [3]
合併run after append [3 5]
合併left_arr [2] right_arr [3 5]
合併run after for [2]
合併run after append [2 3 5]
合併left_arr [1 4 7] right_arr [2 3 5]
合併run after for [1 2 3 4 5]
合併run after append [1 2 3 4 5 7]
拆分 arr [8 10 33 6 77 100 101]
拆分 arr [8 10 33]
拆分 arr [8]
拆分 arr [10 33]
拆分 arr [10]
拆分 arr [33]
合併left_arr [10] right_arr [33]
合併run after for [10]
合併run after append [10 33]
合併left_arr [8] right_arr [10 33]
合併run after for [8]
合併run after append [8 10 33]
拆分 arr [6 77 100 101]
拆分 arr [6 77]
拆分 arr [6]
拆分 arr [77]
合併left_arr [6] right_arr [77]
合併run after for [6]
合併run after append [6 77]
拆分 arr [100 101]
拆分 arr [100]
拆分 arr [101]
合併left_arr [100] right_arr [101]
合併run after for [100]
合併run after append [100 101]
合併left_arr [6 77] right_arr [100 101]
合併run after for [6 77]
合併run after append [6 77 100 101]
合併left_arr [8 10 33] right_arr [6 77 100 101]
合併run after for [6 8 10 33]
合併run after append [6 8 10 33 77 100 101]
合併left_arr [1 2 3 4 5 7] right_arr [6 8 10 33 77 100 101]
合併run after for [1 2 3 4 5 6 7]
合併run after append [1 2 3 4 5 6 7 8 10 33 77 100 101]
[1 2 3 4 5 6 7 8 10 33 77 100 101]
*/