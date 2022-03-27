package main

import "fmt"

func main(){
	var arr = []int{2,7,11,15}
	fmt.Print(twoSum(arr,9))
}

func twoSum(nums []int, target int)[]int{
	m:=make(map[int]int)
	for i := 0; i < len(nums); i++ {
		c,ok:= m[nums[i]]
		/*
			Golang提供一個ok值檢查map裡有沒有該值。
			如果沒有，就把target - 該值當作key存進map，value是index
		*/
		if ok {
			r:=[]int{c,i}
			return r
		}else{
			m[target-nums[i]]=i
		}
	}
	r:=[]int{0,0}
	return r
}

