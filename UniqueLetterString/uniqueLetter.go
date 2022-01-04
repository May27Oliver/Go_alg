package main

import "fmt"

func main() {
	a:=uniqueLetterString("thisishowwedoit")
	fmt.Print(a)
}

//找到一個字串內不重複字串的最長子字串
//技巧：sliding window遍歷字串、counter紀錄字母
func uniqueLetterString(str string) string {
	startPt := 0
	endPt := 0
	counter := map[string]int{}
	var apha, res, subStr string//宣告字母、子字串、結果變數
	maxLen := 0
	for startPt < len(str) {
		if endPt >= len(str){
			subStr = str[startPt:endPt]
			if maxLen < len(subStr){
				maxLen = len(subStr)
				res = subStr
			}
			break
		}

		apha = string(str[endPt])
		counter[apha]++//讓endPt一個字一個字往後走，每走到一個字上就在counter內記錄這個字(key)，出現了幾次(value)

		if counter[apha] > 1 {//出現字母重複！！紀錄子字串、紀錄子字串長度
			subStr = str[startPt:endPt-1]
			tempApha := string(str[startPt])
			
			if maxLen < len(subStr){
				maxLen = len(subStr)
				res = subStr
			}

			if tempApha == apha {
				counter[apha]--
			}else{
				counter[apha]--
				counter[tempApha] = 0
			}
			startPt++
		} else if counter[apha] == 1 {//字母沒有重複，endPt繼續往下走
			endPt++
		}
	}
	return res
}
