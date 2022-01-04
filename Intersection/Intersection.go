package main

import "fmt"

func main(){
	b:=stringIntersec("aabbddc","ababddc")
	fmt.Print(b);
}

func intersect(nums1 []int, nums2 []int) []int {
	counter := map[int]int{}
	var res []int
	for _, n := range nums1 {
		counter[n]++
	}
	for _, n := range nums2 {
		if counter[n] > 0 {

			res = append(res, n)
			counter[n]--
		}
	}
	return res
}

func stringIntersec(str1 string, str2 string)bool{
	if(len(str1)!=len(str2)){//如果兩個字串長度不一樣直接不比
		return false
	}
	counter := map[string]int{}
	var res bool
	for _,v := range str1{
		counter[string(v)]++
	}

	for _,v := range str2{
		if counter[string(v)] > 0 {
			counter[string(v)]++
		}
	}

	for _,v :=range counter{
		if v == 1 {
			res = false
		}else{
			res = true
		}
	}

	return res
}
// package leetcode
