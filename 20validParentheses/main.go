package main

import (
	"fmt"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

func main(){
	a:=isValid(")")
	fmt.Print(a)
}
/*
思路:
stack先進後出
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}
	var stack []string
	left := []string{"(","[","{"}
	right := []string{")","]","}"}
	var m = map[string]string{
		")":"(",
		"}":"{",
		"]":"[",
	}
	for _,v:=range s{
		t:=string(v)
		//左括號放入stack
		if IndexOf(left, t) > -1{
			stack = append(stack,t)
		}
		//右括號
		if IndexOf(right, t) > -1{
			if len(stack) == 0{
				return false
			}
			k:=len(stack)-1
			lastStr :=stack[k]
			stack = stack[:k]
			fmt.Printf("lastStr == %s\n",lastStr)
			fmt.Printf("m[t] == %s\n",m[t])

			if m[t] != lastStr{
				return false
			} 
		}
	}
	return len(stack) == 0
}

func IndexOf(arr []string, s string)int{
	for i,v := range arr {
		if v == s{
			return i
		}
	}
	return -1
}

// error
// func checkValid(s string)bool{
// 	var m = map[string]string{
// 		"(":")",
// 		"{":"}",
// 		"[":"]",
// 	}
// 	length := len(s)
//     if length ==1 || length == 3 || length == 5 {
// 		return false
// 	}
// 	i := strings.Index(s,m[s[:1]])
// 	fmt.Println(i)
// 	switch i{
// 		case length - 1:
// 			if length == 2 {
// 				return true
// 			} 
// 			return checkValid(s[1:length-1])
// 		case 1:
// 			if length == 2 {
// 				return true
// 			} 
// 			return checkValid(s[2:])
// 		default:
// 			return false
// 	}
// }