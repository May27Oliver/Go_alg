package main

import (
	"fmt"
	"strings"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

func main(){
	var s = []string{"c","acc","ccc"}
	a:=longestCommonPrefix(s)
	fmt.Print(a)
}

func longestCommonPrefix(strs []string) string {
	var s = strs[0]
	var substr string 
	for i := 1; i <= len(s); i += 1 {
		substr = s[:i]
		for j := 1; j < len(strs); j += 1 {
			if strings.Index(strs[j],substr) != 0 {
				if len(substr) > 1{
					return s[:i-1]
				}else{
					return ""
				}
			}
		}
	}
	return s
}
