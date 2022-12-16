package main

import (
	"fmt"
)

func main(){

	array := [][]int{
		//{10, 1, 2, 3, 4, 5},
		//{1,3,2},
		{1,3,2,1},
		{1,2,1,2},
		{3,6,5,8,10,20,15},

		}

	for _,x:= range array{
		fmt.Println(almostIncreasingSequence(x))
	}

}

func almostIncreasingSequence(sequence []int) bool {
	isSequence := true
	removedElements :=0
	var existingElements  = map[int]int{}

	for i,j:=1,0; i<len(sequence); i++{
		isMayor := sequence[j] >= sequence[i]
		//res := sequence[i] - sequence[j]
		if _,exist := existingElements[sequence[i]]; exist{
			removedElements++
			j++
			continue
		}
		if isMayor{
			removedElements++
			j++
			existingElements[sequence[j] = 1
			continue
		}
		existingElements[sequence[j]] = 1
		j++
	}
	if removedElements > 1{
		return false
	}

	return isSequence
}


func padespues(sequence []int) bool {
	isSequence := true
	removedElements :=0

	for i,j:=1,0; i<len(sequence); i++{
		res := sequence[i] - sequence[j]
		if res > 1 {
			sequence = append(sequence[0:i],sequence[i+1:]...)
			removedElements++
			j--
			i--
		} else if res < 1 {
			sequence = sequence[i:]
			removedElements++
			j--
			i--
		}
		j++
	}
	if removedElements > 1{
		return false
	}

	return isSequence
}