package main

import "fmt"

//設定兩個pointer，一個遍歷陣列，一個紀錄獨一無二之數
func main() {
	list := []int{1,1,2}
	a:=removeDuplicates(list)
	fmt.Print(a)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	ranger,lord:=0,0
	for ranger < len(nums) {
		if nums[ranger] != nums[lord]{
			lord++
			nums[lord] = nums[ranger]
		}
		ranger++
	}
	return lord+1
}
