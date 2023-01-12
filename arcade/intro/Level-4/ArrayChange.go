package main

import "fmt"

func arrayChange(inputArray []int) int {
	sum := 0
	for i:=1; i<len(inputArray);i++ {
		if dif := inputArray[i]-inputArray[i-1]; dif <= 0{
			inputArray[i] += (dif *-1)+1
			sum += (dif *-1)+1
		}
	}
	return sum
}

func main(){

	fmt.Println( arrayChange([]int{1,1,1} ))

	/*fmt.Println( arrayChange([]int{-1000, 0, -2, 0} ))
	fmt.Println( arrayChange([]int{2, 1, 10, 1} ))
	fmt.Println( arrayChange([]int{2, 3, 3, 5, 5, 5, 4, 12, 12, 10, 15} ))
	fmt.Println( arrayChange([]int{} ))
	fmt.Println( arrayChange([]int{} ))
	fmt.Println( arrayChange([]int{} ))
*/

}
