package main

import "fmt"

func main() {
	a := minSubLen([]int{8, 1, 6, 15, 3, 16, 5, 7, 14, 30, 12}, 60)
	fmt.Print(a)
}

func minSubLen(arr []int, num int) int {
	//思維：先找陣列中最大值，然後slide window，看最小長度幾大於指定數
	//在找陣列次大值....
	startPt := 0
	endPt := 0
	minLen := len(arr)
	temp_sum := arr[startPt]
	for startPt < len(arr) {
		if temp_sum < num && endPt < len(arr) {
			temp_sum += arr[endPt]
			endPt++
		} else if temp_sum >= num {
			currentLen := endPt - startPt
			if minLen > currentLen {
				minLen = currentLen
			}
			temp_sum -= arr[startPt]
			startPt++
		} else if endPt >= len(arr) {
			break
		}
	}
	if minLen == 0 {
		fmt.Print("Cannot find MinLen")
		return -1
	}
	return minLen
}
