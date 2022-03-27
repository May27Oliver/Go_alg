package main

import (
	"fmt"
	"strconv"
)


func main(){
	s:="111"
	a:=Solution(s)
	fmt.Print(a)
}

func Solution(S string) int {
	// write your code in Go 1.4
	//將二進位string轉為十進位
	num, _:= strconv.ParseInt(S, 2, 0)
	V,count := int(num),1
	for V > 1{
		if V % 2 == 0{//奇數減一，偶數除二，直到V == 0,用count紀錄次數
			V = V / 2
		}else{
			V = V - 1
		}
		count++
	}
	return count
}
