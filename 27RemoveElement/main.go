package main

import "fmt"

func main(){
	a:=[]int{2,3,2,4,3,5,2}
	b:=removeElement(a, 2)
	fmt.Print(b)
}

func removeElement(nums []int, val int) int {
	//快慢指針
	ranger,lord:=0,0
	for ranger < len(nums){
		if nums[ranger] != val {
			nums[lord] = nums[ranger]
			lord++
		}
		ranger++
	}
	return lord
}
// func removeElement(nums []int, val int) int {
//     /*
// 			[2,3,2,1,5,6,7,3,6]
// 			在nums陣列中把值等於val的值給刪除掉，並將nums改成沒有空值的陣列。
// 		*/
// 		if nums == nil{
// 			return 0
// 		}
// 		for i:=0;i<len(nums);{
// 			if nums[i] == val{
// 				nums = append(nums[:i],nums[i+1:]...)
// 				i--
// 			}
// 			i++
// 		}
// 		return len(nums)
// } 