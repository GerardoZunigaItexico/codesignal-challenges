package main

import "fmt"

func palindromeRearranging(inputString string) bool {
	ris := []rune(inputString)
	m :=make(map[rune]int)
	oddC := 0
	for _,v:= range ris {
		if vv,ok := m[v]; ok{
			m[v] = vv +1
		}else{
			m[v]=1
		}
	}
	for _,v := range m{
		if v%2!=0{
			oddC++
		}
		if oddC>1{
			return false
		}
	}

	return true
}


func reverse(inputRune []rune) string{
	l := len(inputRune)
	for i := l/2-1; i >= 0; i-- {
		opp := l-1-i
		inputRune[i], inputRune[opp] = inputRune[opp], inputRune[i]
	}
	return string(inputRune)
}

func main(){

	fmt.Println( palindromeRearranging("aabb" ))
	fmt.Println( palindromeRearranging("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc" ))


}
