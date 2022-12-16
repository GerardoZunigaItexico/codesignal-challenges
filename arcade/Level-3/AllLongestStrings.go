package main

import (
	"fmt"
)

func allLongestStrings(inputArray []string) []string {
	var outputArray []string
	l := 0
	for _,v := range inputArray {
		nL := len(v)
		if nL>l {
			outputArray = []string{}
			l=nL
		}
		if nL<l{
			continue
		}
		outputArray = append(outputArray,v)

	}
	return outputArray
}


func main(){
	inputArray := []string{"aba", "aa", "ad", "vcd", "aba"}

	fmt.Println(allLongestStrings(inputArray))
}

