package main

import "fmt"

func main(){
	a := MaxAndMinSum([]int{2,7,3,0,6,1,-5,-12,-11},2)
	fmt.Print(a)
}

//refractor
func MaxAndMinSum(arr []int,num int)[]int{
	if num > len(arr){
		return []int{}
	}
	var temp_sum int
	var max int
	var min int 
	for a:= 0 ; a < num; a++ {
		max += arr[a]
	}
	min = max
	temp_sum = max
	for b:= num; b < len(arr); b++ {
		temp_sum = temp_sum + arr[b] - arr[ b - num ]
		if temp_sum > max {
			max = temp_sum
		}
		if temp_sum < min {
			min = temp_sum
		}
	}
	return []int{max,min}
}