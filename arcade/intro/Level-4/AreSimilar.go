package main

import "fmt"

func areSimilar(a []int, b []int) bool {
	swapped := false
	for i:=range a{
		if a[i] == b[i]{
			continue
		}
		if swapped {
			return false
		}
		for j:=i; j<len(a);j++{
			if a[i] == b[j] && a[j] != b[j]{
				b[i], b[j] =b[j],b[i]
				swapped=true
				break
			}
		}
		if !swapped{
			return false
		}
	}
	return true
}

func main(){

	fmt.Println( areSimilar([]int{1,2,3},[]int{1,2,3}) )
	fmt.Println( areSimilar([]int{1,2,3},[]int{2,1,3}) )
	fmt.Println( areSimilar([]int{1,2,2},[]int{2,1,1}) )
	fmt.Println( areSimilar([]int{1,1,4},[]int{1,2,3}) )
	fmt.Println( areSimilar([]int{2,3,1},[]int{1,3,2}) )
	fmt.Println( areSimilar([]int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},[]int{832, 998, 148, 570, 533, 561, 455, 147, 894, 279}) )
}
