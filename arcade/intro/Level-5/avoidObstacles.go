package main

import (
	"fmt"
	"sort"
)

func avoidObstacles(inputArray []int) int {
	sort.Ints(inputArray)
	lia := len(inputArray)
	mJump := []int{}
	for i,lia:=1,lia; i<lia;i++{
		if r:=inputArray[i] - inputArray[i-1]; r>1{
			mJump = append(mJump, inputArray[i-1]+1)
		}
	}
	lmj := len(mJump)
	if  lmj == 0 {
		return inputArray[lia-1]+1
	}
	if lmj == 1 {
		return mJump[0]
	}
	maxValid := 0
	for i:=lmj-2;lmj>1&&i>=0;i--{
		if newMax := mJump[i+1]-mJump[i]; newMax > maxValid{
			maxValid =  newMax
		}
	}

	return maxValid
}

func main(){
	//fmt.Println(avoidObstacles([]int{5, 3, 6, 7, 9}))
	//fmt.Println(avoidObstacles([]int{2,3}))
	fmt.Println(avoidObstacles([]int{1, 4, 10, 6}))
}
