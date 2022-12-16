package main

import (
	"fmt"
)

func main(){
	fmt.Println(checkPalindrome("aabaa"))
	fmt.Println(checkPalindrome("abac"))

}

func checkPalindrome(inputString string) bool {
	var noSpace string
	for _, v := range inputString{
		value := string(v)
		noSpace += value
	}

	for i,j:=0,len(noSpace)-1; i<(len(noSpace)/2)+1; i++{
		if noSpace[i] != noSpace[j]{
			return false
		}

		if ((len(noSpace)%2)==1) && i == j{
			return true
		}else if ((len(noSpace)%2)==0) && i+1 == j{
			return true
		}
		j--
	}
	return false
}



