package main

import "fmt"

func main(){
	ans:=isPalindrome(121)
	fmt.Print(ans)
}
//思路，把數字從個位數一個一個取下，變成另外一個迴文的數字
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origin := x
	var reverse = 0
	for origin > 0{
		v := origin % 10
		origin = origin / 10
		reverse = reverse*10 + v
	}
	return reverse == x
}