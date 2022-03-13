package main

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main(){
	a:=Solution("zyzyzyz")
	fmt.Print(a)
}

func Solution(S string) int {
	if S == ""{
			return 0
	}
	//sliding window and counter   
	i,winLen := 0,2  //起始pointer與substring長度
	var counter = map[string]int{} //紀錄版

	for winLen < len(S){
			for i < len(S) - winLen + 1{ //sliding winlen長度的子字串在整個字串
					counter[S[i: i + winLen]]++
					i++
			}
			//取出key值，看是否有substring重複數為一不重複
			for k := range counter {
					if counter[k] == 1{
							fmt.Printf("substring %s is unique",k)
							return winLen
					}
			}
			winLen++
			i = 0
	} 
	return 0
}