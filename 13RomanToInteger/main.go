package main

import (
	"fmt"
)

var (
	I = 1 
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)
/*
I = 1
II = 2
III = 3
IV = 4
V = 5
VI = 6
VII = 7
VIII = 8
IX = 9
X = 10
IC = 99

Given a roman numeral, convert it to an integer.

strings.Index(s)int 查詢substring在strings中的索引，-1代表不包含。
strings.LastIndex(s)int 查詢substring在strings中最後出現的索引，-1代表不包含。
往右是加
往左是減
先找到字串中數字最大的字母，最該字母往左是減，往右是加

在golang中，字串的遍歷不能直接取索引值，會印出utf-8編碼後的位元組，要將字串轉換成[]rune
*/
func main(){
	a:=romanToInt("MDCLV")
	fmt.Print(a)
}
//way1.
// func romanToInt(s string) int {
// 	var list = map[string]int{
// 		"M":1000,
// 		"D":500,
// 		"C":100,
// 		"L":50,
// 		"X":10,
// 		"V":5,
// 		"I":1,
// 	}
// 	rs := []rune(s)//將字串轉為unicode
// 	//設兩個pointer由左到右比較字串，左小於右減，否則兩兩加上
// 	pointer,temp, sum := 0,0,0
// 	for pointer < len(rs) {
// 		cur_val := list[string(rs[pointer])]
// 		if cur_val > temp {
// 			sum = sum + cur_val -  temp * 2
// 		}else{
// 			sum += cur_val
// 		}		
// 		temp = list[string(rs[pointer])]
// 		pointer++
// 	}
// 	return sum
// }

//way2
// func romanToInt(s string) int {
// 	sum:=0
// 	var list = map[string]int{
// 		"M":1000,
// 		"D":500,
// 		"C":100,
// 		"L":50,
// 		"X":10,
// 		"V":5,
// 		"I":1,
// 	}
// 	last:=0
// 	for i := len(s)-1 ; i >= 0; i-- {
// 		temp:=list[string(s[i])]
// 		sign:=1
// 		if temp < last{
// 			sign = -1
// 		}
// 		sum += sign*temp
// 		last = temp
// 	}	
// 	return sum
// }

func romanToInt(s string) int {
	sum:=0
	var list = map[string]int{
		"M":1000,
		"D":500,
		"C":100,
		"L":50,
		"X":10,
		"V":5,
		"I":1,
	}
	last:=0
	for i := len(s)-1; i >= 0; i -= 1 {
		temp := list[string(s[i])] 
		sign:=1
		if last > temp{
			sign = -1
		}
		sum += sign*temp
		last = temp
	}
	return sum
}