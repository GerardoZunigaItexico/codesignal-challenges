package main

import (
	"fmt"
	"strings"
)



func reverseInParentheses(inputString string) string {
	openedP, closedP := fetchParenthesesPositions(inputString)
	if openedP==0 && closedP == 0{
		return inputString
	}

	word := inputString[ openedP+1:closedP ]
	reversed := reverse(word)
	inputString = strings.Replace(inputString,("("+word+")"),reversed,1)

	return reverseInParentheses(inputString)
}

func reverse(inputString string) string{
	a := []rune(inputString)
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
	return string(a)
}

func fetchParenthesesPositions(inputString string) (int,int){
	lis := len(inputString)
	if lis <= 1{
		return 0,0
	}
	var openedParentheses int
	var closedParentheses int
	for i:=0;i< lis;i++ {
		if string(inputString[i]) == "("{
			openedParentheses = i
		}
		if string(inputString[i]) == ")"{
			closedParentheses = i
		}
		if closedParentheses !=0 && (openedParentheses < closedParentheses){
			break
		}
	}
	return openedParentheses,closedParentheses
}

func main(){
	fmt.Println(reverseInParentheses("foo(bar)baz(blim)"))
	fmt.Println(reverseInParentheses("bar"))
	fmt.Println(reverseInParentheses("foo(bar(baz))blim"))
	fmt.Println(reverseInParentheses("(abc)d(efg)"))
}
