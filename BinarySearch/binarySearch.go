package main

import "fmt"

// import "fmt"

func main(){
	arr := []int{1,2,3,9,16,27,33,89,96,123,189,255,333}
	index:=binarySearch(arr,255)
	fmt.Print(index)
	// a:=math.Floor((0+12)/2)
	// fmt.Println(a);
}

func binarySearch(arr []int, num int) int {
	min := 0
	max := len(arr)-1
	for min <= max{ //Golang while迴圈用 go
		middle := (min + max)/ 2
		if( num > arr[middle]){
			min = middle + 1
		}else if(num < arr[middle]){
			max = middle - 1
		}else if(num == arr[middle]){
			return middle
		}
	}
	return -1
}