package main

import "fmt"

/*
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
*/
func main(){
	a:=countBits(7)
	fmt.Print(a)
}
//我的解法O(n*sizeof(integer))：
// func countBits(n int) []int {
// 	var list []int
//  for i := 0; i <= n ; i += 1 {
// 	list = append(list,total(i))
//  }
//  return list   
// }

// func total(c int)int{
// 	if c == 0 {
// 		return 0 
// 	}
// 	return c % 2 + total(c / 2)
// }

// func countBiOneInHex(n int) int {
// 	countOne:=0
// 	switch n {
// 		case 0:
// 			countOne = 0
// 		case 1,2,4,8:
// 			countOne = 1
// 		case 3,5,6,9,10,12:
// 			countOne = 2
// 		case 7,11,13,14:
// 			countOne = 3
// 		case 15:
// 			countOne = 4
// 	}
// 	return countOne
// }

// func countBinaryOne(n int)int{
// 	r:=0
// 	countOne:=0
// 	for n > 0 {
// 		n, r = n >> 4, n % 16 
// 		countOne += countBiOneInHex(r)
// 	}
// 	return countOne
// }

// func countBits(n int) []int {
// 	s:= make([]int ,n+1)
// 	for i:=0;i<=n;i++{
// 		s[i] = countBinaryOne(i)
// 	}
// 	return s
// }
func countOneTimes(r int) int {
	var n int
	switch r {
		case 1,2,4,8:
			n = 1
		case 3,5,6,9,10,12:
			n = 2
		case 7,11,13,14:
			n = 3
		case 15:
			n = 4	
	}
	return n
}

func countBinaryOne(n int)int{
	//把i化作二進位數
	var r int
	sum := 0
	for n > 0 {
		n,r = n>>4,n%16
		sum += countOneTimes(r)
	}
	return sum	
}

func countBits(n int)[]int{
	var list []int
	s := make([]int,n+1)
	for i := 0; i <= n; i += 1 {
		s[i] = countBinaryOne(i)
	}

	return list
}