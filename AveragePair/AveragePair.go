package main

import "fmt"

func main() { // a:=averagePair([]float32{-11,0,1,2,3,9,14,17,21},1.5)
	a := subsequence("hello", "hello Dear")
	fmt.Printf("hello is hello Dear %d ", a)
}

func averagePair(arr []float32, num float32) [][]float32 {
	pointStart := 0
	pointEnd := len(arr) - 1
	res := [][]float32{}
	for pointEnd > pointStart {
		temp_avg := (arr[pointStart] + arr[pointEnd]) / 2
		if temp_avg > num {
			pointEnd--
		} else if temp_avg < num {
			pointStart++
		} else if temp_avg == num {
			res = append(res, []float32{arr[pointStart], arr[pointEnd]})
			pointStart++
			pointEnd--
		}
	}
	return res
}

func palindrome(str string) bool {
	p1 := 0
	p2 := len(str) - 1
	var res bool = false

	for p2 >= p1 {
		if str[p1] != str[p2] {
			res = false
			break
		} else if str[p1] == str[p2] {
			res = true
			p1++
			p2--
		}
	}
	return res
}

func subsequence(str1 string, str2 string) bool { //比較字串1是不是字串2的subsequence
	st1Pt := 0
	st2Pt := 0
	var res bool = false
	for st2Pt <= len(str2)-1 && st1Pt <= len(str1)-1 {
		if str1[st1Pt] == str2[st2Pt] {
			res = true
			st1Pt++
		} else if str1[st1Pt] != str2[st2Pt] {
			res = false
		}
		st2Pt++
	}
	return res
}
