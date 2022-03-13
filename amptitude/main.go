package main

import "fmt"

func main(){
	a:=[]int{5,3,6,1,3}
	s:=Solution(a,2)
	fmt.Println(s)
}

func Solution(A []int, K int) int {
	// write your code in Go 1.4
	//刪去K個值，讓陣列裡的差距最小
	//先把陣列按順序大小排好
	arr:= mergeSort(A)
	fmt.Print(arr)
	return 1
}
func mergeSort(arr []int) []int {
	if len(arr) < 2{
		return arr
	}
	//拆分
	midIdx := len(arr) / 2
	arr_left := mergeSort(arr[:midIdx])
	arr_right := mergeSort(arr[midIdx:])
	//合併
	return merge(arr_left, arr_right)
}

func merge(left_arr, right_arr []int) []int {
	// fmt.Println("合併left_arr", left_arr, "right_arr", right_arr)
	final := []int{}
	i := 0
	j := 0
	for i < len(left_arr) && j < len(right_arr) {
		if left_arr[i] < right_arr[j] {
			final = append(final, left_arr[i])
			i++
		} else {
			final = append(final, right_arr[j])
			j++
		}
	}
	//上面迴圈跑完，左右陣列一定還會遺留一些值，大於final內的任何值，下面用loop將之加入
	// fmt.Println("合併run after for", final)
	for ; i < len(left_arr); i++ {
		final = append(final, left_arr[i])
	}
	for ; j < len(right_arr); j++ {
		final = append(final, right_arr[j])
	}
	// fmt.Println("合併run after append", final)
	return final
}